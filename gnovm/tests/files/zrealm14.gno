// PKGPATH: gno.land/r/test
package test

var (
	a = &A{A: "here"}
	b [2]*B
)


type A struct {
	A string
}
type B struct {
	A *A
	B string
}

func init() {
 	c := B{
		A: a,
		B: "c",
	}
	b[0] = &c

	d := B{
		A: a,
		B: "d",
	}
	b[1] = &d
}

func main() {
	b[0] = nil
	b[1] = nil
}

// Realm:
// switchrealm["gno.land/r/test"]
// u[a8ada09dee16d791fd406d629fe29bb0ed084a30:5]={
//     "Data": null,
//     "List": [
//         {
//             "T": {
//                 "@type": "/gno.PointerType",
//                 "Elt": {
//                     "@type": "/gno.RefType",
//                     "ID": "gno.land/r/test.B"
//                 }
//             }
//         },
//         {
//             "T": {
//                 "@type": "/gno.PointerType",
//                 "Elt": {
//                     "@type": "/gno.RefType",
//                     "ID": "gno.land/r/test.B"
//                 }
//             }
//         }
//     ],
//     "ObjectInfo": {
//         "ID": "a8ada09dee16d791fd406d629fe29bb0ed084a30:5",
//         "ModTime": "9",
//         "OwnerID": "a8ada09dee16d791fd406d629fe29bb0ed084a30:2",
//         "RefCount": "1"
//     }
// }
// u[a8ada09dee16d791fd406d629fe29bb0ed084a30:3]={
//     "ObjectInfo": {
//         "ID": "a8ada09dee16d791fd406d629fe29bb0ed084a30:3",
//         "IsEscaped": true,
//         "ModTime": "9",
//         "RefCount": "1"
//     },
//     "Value": {
//         "T": {
//             "@type": "/gno.RefType",
//             "ID": "gno.land/r/test.A"
//         },
//         "V": {
//             "@type": "/gno.RefValue",
//             "Hash": "95dca2f5b12899b6367402ecdac04c7ca59a03d9",
//             "ObjectID": "a8ada09dee16d791fd406d629fe29bb0ed084a30:4"
//         }
//     }
// }
// d[a8ada09dee16d791fd406d629fe29bb0ed084a30:8]
// d[a8ada09dee16d791fd406d629fe29bb0ed084a30:9]
