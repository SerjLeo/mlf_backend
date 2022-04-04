package password

import "testing"

const testSalt = "hashTest123"

var testCases = []struct {
	in  string
	out string
}{
	{"fdka874321rldsavn", "686173685465737431323397e9c27a2f46139d6b507afd96f2a133d2dbb4a7"},
	{"fdal;0fdas0fa34", "6861736854657374313233c15c6c72b9aea878a48c7d655cd0af6d5ec56626"},
	{"gfs.h/fdgrte", "6861736854657374313233649972dc9d03ab09dcb020c15f162eaddbbc96e5"},
	{"53428fdsam9423az", "686173685465737431323361b1dc476d86baa949f819982a5764765dba8c8d"},
	{"аыдgkfdдпжав43@fdsv--0", "68617368546573743132331dc99f35f90497fbe7cf5fc1f9d60ee31ddd1d1c"},
}

func TestSHA1Hash_EncodeString(t *testing.T) {
	hashService := NewSHA1Hash(testSalt)

	t.Run("testing success cases", func(t *testing.T) {
		for _, tt := range testCases {
			hash, err := hashService.EncodeString(tt.in)
			if err != nil {
				t.Errorf("Error while hashing, %s", err.Error())
			}
			if hash != tt.out {
				t.Errorf("got %q, want %q", hash, tt.out)
			}
		}
	})

	anotherHashService := NewSHA1Hash("wrongSalt")

	t.Run("testing fail cases (changed salt)", func(t *testing.T) {
		for _, tt := range testCases {
			hash, err := anotherHashService.EncodeString(tt.in)
			if err != nil {
				t.Errorf("Error while hashing, %s", err.Error())
			}
			if hash == tt.out {
				t.Errorf("got %q and %q equal, should be different", hash, tt.out)
			}
		}
	})

}
