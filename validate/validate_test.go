package validate

import(
	testing
)

func TestFileInput1(t *testing.T){
	filename:="test.htm"
	want:="regexp.MustCompile"
	msg,err = checkFilename()
}