package aefire

type AEError string

func (e AEError) Error() string {
	return string(e)
}

const (
	DocumentNotFound = AEError("DocumentNotFound")
)

/*func NewBatch() *firestore.WriteBatch {
	return FStore.Batch()
}

func Doc(path string) *firestore.DocumentRef {
	return FStore.Doc(path)
}

func Snap(path string) *firestore.DocumentSnapshot {
	snap, err := Doc(path).Get(context.Background())

	if status.Code(err) == codes.NotFound {
		return nil
	} else {
		return snap
	}
}

func Col(path string) *firestore.CollectionRef {
	return FStore.Collection(path)
}

func SnapExists(snap *firestore.DocumentSnapshot) bool {
	return snap != nil && snap.Exists()
}

func SnapTo(snap *firestore.DocumentSnapshot, to interface{}) error {
	if !SnapExists(snap) {
		return DocumentNotFound
	}

	return snap.DataTo(to)
}

func PathTo(path string, to interface{}) error {
	snap := Snap(path)
	if !SnapExists(snap) {
		return DocumentNotFound
	}

	return snap.DataTo(to)
}

func QueryStringField(q firestore.Query, field string) (fieldValues map[string]string) {
	iter := q.Documents(context.Background())

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

func BatchMerge(b *firestore.WriteBatch, doc *firestore.DocumentRef, pairs ...interface{}) *firestore.WriteBatch {
	b.Set(doc, MapOf(pairs...), firestore.MergeAll)

	return b
}

func BatchMergePath(b *firestore.WriteBatch, path string, pairs ...interface{}) *firestore.WriteBatch {
	b.Set(Doc(path), MapOf(pairs...), firestore.MergeAll)

	return b
}

func BatchMergeObj(b *firestore.WriteBatch, path string, o interface{}) *firestore.WriteBatch {
	b.Set(Doc(path), ToMap(o), firestore.MergeAll)

	return b
}

func DocMerge(doc *firestore.DocumentRef, pairs ...interface{}) error {
	_, err := doc.Set(context.Background(), MapOf(pairs...), firestore.MergeAll)
	return err
}

func PathMerge(path string, pairs ...interface{}) error {
	return DocMerge(FStore.Doc(path), pairs...)
}

func DocMergeObj(doc *firestore.DocumentRef, o interface{}) error {
	_, err := doc.Set(context.Background(), ToMap(o), firestore.MergeAll)
	return err
}

func DeleteQueryResults(q firestore.Query, b *firestore.WriteBatch) error {
	snaps, err := q.Documents(context.Background()).GetAll()
	if err != nil {
		return err
	}

	newBatch := false

	if b == nil {
		b = NewBatch()
		newBatch = true
	}

	for _, snap := range snaps {
		b.Delete(snap.Ref)
	}

	if !newBatch {
		return nil
	} else {
		_, err = b.Commit(context.Background())
		return err
	}
}
func UpdateQueryResults(q firestore.Query, b *firestore.WriteBatch, updates ...interface{}) error {
	snaps, err := q.Documents(context.Background()).GetAll()
	if err != nil {
		return err
	}

	newBatch := false
	if b == nil {
		b = NewBatch()
		newBatch = true
	}

	for _, snap := range snaps {
		BatchMerge(b, snap.Ref, updates...)
	}

	if !newBatch {
		return nil
	} else {
		_, err = b.Commit(context.Background())
		return err
	}
}

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
