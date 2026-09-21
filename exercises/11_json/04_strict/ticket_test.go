package ticket

import "testing"

func TestUnmarshalStrict(t *testing.T) {
	got, err := UnmarshalStrict(`{"id":1,"title":"t","status":"open"}`)
	if err != nil || got.Title != "t" {
		t.Fatalf("valid JSON = (%+v, %v)", got, err)
	}
	if _, err := UnmarshalStrict(`{"id":1,"title":"t","zzz":1}`); err == nil {
		t.Fatal("unknown field accepted, want strict error")
	}
}
