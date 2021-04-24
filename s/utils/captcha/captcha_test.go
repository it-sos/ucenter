package captcha

import "testing"

func TestGenerate(t *testing.T) {
	tests := []struct {
		name     string
		wantId   string
		wantB64s string
		wantErr  bool
	}{
		{
			"name",
			"1",
			"?",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotB64s := Generate()
			t.Log(gotId, gotB64s, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("Generate() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotB64s != tt.wantB64s {
				t.Errorf("Generate() gotB64s = %v, want %v", gotB64s, tt.wantB64s)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Verify("", ""); got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}
