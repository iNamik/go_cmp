package cmp

// Compare Results
const (
	LT int = -1 // Less-Than
	EQ     = 0  // Equal
	GT     = 1  // Greater-Than
)

// T allows objects to export a comparator function
type T interface {
	Cmp(interface{}) int
}

// F allows you to pass functions that perform comparisons
type F func(a, b interface{}) int

// TF allows you to cast a func/closure into cmp.T interface.
// Example:
//	var key cmp.T = cmp.TF(func(v interface{}) int { /*...*/ })
type TF func(interface{}) int

// TF::Cmp fulfills the cmp.T interface.
// NOTE: Yes you can treat functions as interfaces. Per the FAQ, "Methods can be implemented for any type"
//       http://golang.org/doc/faq#principles
//       http://jordanorelli.tumblr.com/post/42369331748/function-types-in-go-golang
func (f TF) Cmp(v interface{}) int {
	return f(v)
}

// NewTF takes a comp.F and an initial value, 'a',
// and creates partially-applied function of type cmp.TF (fulfilling cmp.T).
// The resulting cmp.TF/cmp.T can be compared to multiple 'b' values.
// Example:
//  func find(needle_ interface{}, hayStack []interface{}, fcmp cmp.F) (int, bool){
//    var needle cmp.T = cmp.NewTF(fcmp, needle_)
//    for i, hay := range hayStack {
//      if needle.Cmp(hay) == cmp.EQ {
//	      return i, true
//      }
//    }
//    return -1, false
//  }
func NewTF(f F, a interface{}) TF {
	return TF(func(b interface{}) int { return f(a, b) })
}
