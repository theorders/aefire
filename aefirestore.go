package aefire

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (a *AEFire) Doc(path ...string) *firestore.DocumentRef {
	return a.FStore.Doc(strings.Join(path, "/"))
}

func (a *AEFire) Snap(path ...string) (*firestore.DocumentSnapshot, error) {
	snap, err := a.Doc(path...).Get(a)

	if status.Code(err) == codes.NotFound {
		return nil, DocumentNotFound
	} else {
		return snap, err
	}
}

func (a *AEFire) DocSnap(doc *firestore.DocumentRef) (*firestore.DocumentSnapshot, error) {
	snap, err := doc.Get(a)

	if status.Code(err) == codes.NotFound {
		return nil, DocumentNotFound
	} else {
		return snap, err
	}
}

func (a *AEFire) Col(path ...string) *firestore.CollectionRef {
	return a.FStore.Collection(strings.Join(path, "/"))
}

func (a *AEFire) SnapTo(snap *firestore.DocumentSnapshot, to interface{}) error {
	if snap == nil || snap.Exists() == false {
		return DocumentNotFound
	}

	return snap.DataTo(to)
}

func SnapTo(snap *firestore.DocumentSnapshot, to interface{}) error {
	if snap == nil || snap.Exists() == false {
		return DocumentNotFound
	}

	return snap.DataTo(to)
}

func (a *AEFire) PathTo(to interface{}, path ...string) error {
	snap, err := a.Snap(path...)
	if err != nil {
		return err
	}

	return snap.DataTo(to)
}

func (a *AEFire) DocTo(to interface{}, ref *firestore.DocumentRef) error {
	snap, err := a.DocSnap(ref)
	if err != nil {
		return err
	}

	return snap.DataTo(to)
}

func (a *AEFire) QueryStringField(q firestore.Query, field string) (fieldValues map[string]string) {
	fieldValues = map[string]string{}
	iter := q.Documents(a)

	for {
		snap, err := iter.Next()
		if err == iterator.Done {
			break
		}

		value, err := snap.DataAt(field)
		if str, ok := value.(string); ok {
			fieldValues[snap.Ref.ID] = str
		}
	}

	return
}

func (a *AEFire) BatchMerge(b *firestore.WriteBatch, doc *firestore.DocumentRef, pairs ...interface{}) *firestore.WriteBatch {
	b.Set(doc, MapOf(pairs...), firestore.MergeAll)

	return b
}

func BatchMerge(b *firestore.WriteBatch, doc *firestore.DocumentRef, pairs ...interface{}) *firestore.WriteBatch {
	b.Set(doc, MapOf(pairs...), firestore.MergeAll)

	return b
}

func (a *AEFire) BatchMergePath(b *firestore.WriteBatch, path string, pairs ...interface{}) *firestore.WriteBatch {
	b.Set(a.Doc(path), MapOf(pairs...), firestore.MergeAll)

	return b
}

func (a *AEFire) BatchMergeObj(b *firestore.WriteBatch, path string, o interface{}) *firestore.WriteBatch {
	b.Set(a.Doc(path), ToMap(o), firestore.MergeAll)

	return b
}

func (a *AEFire) DocMerge(doc *firestore.DocumentRef, pairs ...interface{}) error {
	_, err := doc.Set(a, MapOf(pairs...), firestore.MergeAll)
	return err
}

func (a *AEFire) PathMerge(path string, pairs ...interface{}) error {
	return a.DocMerge(a.FStore.Doc(path), pairs...)
}

func (a *AEFire) DocMergeObj(doc *firestore.DocumentRef, o interface{}) error {
	_, err := doc.Set(a, ToMap(o), firestore.MergeAll)
	return err
}

func (a *AEFire) DeleteQueryResults(q firestore.Query, b *firestore.WriteBatch) error {
	snaps, err := q.Documents(a).GetAll()
	if err != nil {
		return err
	}

	newBatch := false

	if b == nil {
		b = a.FStore.Batch()
		newBatch = true
	}

	for _, snap := range snaps {
		b.Delete(snap.Ref)
	}

	if !newBatch {
		return nil
	} else {
		_, err = b.Commit(a)
		return err
	}
}
func (a *AEFire) UpdateQueryResults(q firestore.Query, b *firestore.WriteBatch, updates ...interface{}) error {
	snaps, err := q.Documents(a).GetAll()
	if err != nil {
		return err
	}

	if len(snaps) == 0 {
		return nil
	}

	newBatch := false
	if b == nil {
		b = a.FStore.Batch()
		newBatch = true
	}

	for _, snap := range snaps {
		a.BatchMerge(b, snap.Ref, updates...)
	}

	if !newBatch {
		return nil
	} else {
		_, err = b.Commit(a)
		return err
	}
}

/*
func UpdatesOf(pairs ...interface{}) []firestore.Update {
	if len(pairs)%2 != 0 {
		panic("MapOf: key-value pair cannot be odd")
	}

	u := []firestore.Update{}

	for i, kv := range pairs {
		if i%2 == 1 {
			k := pairs[i-1].(string)
			u = append(u, firestore.Update{
				Path:  k,
				Value: kv,
			})
		}
	}

	return u
}
*/
