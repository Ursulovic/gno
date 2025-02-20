// Code generated by the genstd tool (@/misc/genstd); DO NOT EDIT.
// To regenerate it, run `go generate` from this directory.

package stdlibs

import (
	"reflect"

	gno "github.com/gnolang/gno/gnovm/pkg/gnolang"
	libs_crypto_ed25519 "github.com/gnolang/gno/gnovm/stdlibs/crypto/ed25519"
	libs_crypto_sha256 "github.com/gnolang/gno/gnovm/stdlibs/crypto/sha256"
	libs_math "github.com/gnolang/gno/gnovm/stdlibs/math"
	libs_std "github.com/gnolang/gno/gnovm/stdlibs/std"
	libs_testing "github.com/gnolang/gno/gnovm/stdlibs/testing"
	libs_time "github.com/gnolang/gno/gnovm/stdlibs/time"
)

// NativeFunc represents a function in the standard library which has a native
// (go-based) implementation, commonly referred to as a "native binding".
type NativeFunc struct {
	gnoPkg     string
	gnoFunc    gno.Name
	params     []gno.FieldTypeExpr
	results    []gno.FieldTypeExpr
	hasMachine bool
	f          func(m *gno.Machine)
}

// HasMachineParam returns whether the given native binding has a machine parameter.
// This means that the Go version of this function expects a *gno.Machine
// as its first parameter.
func (n *NativeFunc) HasMachineParam() bool {
	return n.hasMachine
}

var nativeFuncs = [...]NativeFunc{
	{
		"crypto/ed25519",
		"verify",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("[]byte")},
			{Name: gno.N("p1"), Type: gno.X("[]byte")},
			{Name: gno.N("p2"), Type: gno.X("[]byte")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("bool")},
		},
		false,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  []byte
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  []byte
				rp1 = reflect.ValueOf(&p1).Elem()
				p2  []byte
				rp2 = reflect.ValueOf(&p2).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 2, "")).TV, rp2)

			r0 := libs_crypto_ed25519.X_verify(p0, p1, p2)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"crypto/sha256",
		"sum256",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("[]byte")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("[32]byte")},
		},
		false,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  []byte
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_crypto_sha256.X_sum256(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"math",
		"Float32bits",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("float32")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("uint32")},
		},
		false,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  float32
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_math.Float32bits(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"math",
		"Float32frombits",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint32")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("float32")},
		},
		false,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint32
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_math.Float32frombits(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"math",
		"Float64bits",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("float64")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("uint64")},
		},
		false,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  float64
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_math.Float64bits(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"math",
		"Float64frombits",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint64")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("float64")},
		},
		false,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint64
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_math.Float64frombits(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"bankerGetCoins",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint8")},
			{Name: gno.N("p1"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("[]string")},
			{Name: gno.N("r1"), Type: gno.X("[]int64")},
		},
		true,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint8
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  string
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			r0, r1 := libs_std.X_bankerGetCoins(
				m,
				p0, p1)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
		},
	},
	{
		"std",
		"bankerSendCoins",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint8")},
			{Name: gno.N("p1"), Type: gno.X("string")},
			{Name: gno.N("p2"), Type: gno.X("string")},
			{Name: gno.N("p3"), Type: gno.X("[]string")},
			{Name: gno.N("p4"), Type: gno.X("[]int64")},
		},
		[]gno.FieldTypeExpr{},
		true,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint8
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  string
				rp1 = reflect.ValueOf(&p1).Elem()
				p2  string
				rp2 = reflect.ValueOf(&p2).Elem()
				p3  []string
				rp3 = reflect.ValueOf(&p3).Elem()
				p4  []int64
				rp4 = reflect.ValueOf(&p4).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 2, "")).TV, rp2)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 3, "")).TV, rp3)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 4, "")).TV, rp4)

			libs_std.X_bankerSendCoins(
				m,
				p0, p1, p2, p3, p4)
		},
	},
	{
		"std",
		"bankerTotalCoin",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint8")},
			{Name: gno.N("p1"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("int64")},
		},
		true,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint8
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  string
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			r0 := libs_std.X_bankerTotalCoin(
				m,
				p0, p1)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"bankerIssueCoin",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint8")},
			{Name: gno.N("p1"), Type: gno.X("string")},
			{Name: gno.N("p2"), Type: gno.X("string")},
			{Name: gno.N("p3"), Type: gno.X("int64")},
		},
		[]gno.FieldTypeExpr{},
		true,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint8
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  string
				rp1 = reflect.ValueOf(&p1).Elem()
				p2  string
				rp2 = reflect.ValueOf(&p2).Elem()
				p3  int64
				rp3 = reflect.ValueOf(&p3).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 2, "")).TV, rp2)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 3, "")).TV, rp3)

			libs_std.X_bankerIssueCoin(
				m,
				p0, p1, p2, p3)
		},
	},
	{
		"std",
		"bankerRemoveCoin",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint8")},
			{Name: gno.N("p1"), Type: gno.X("string")},
			{Name: gno.N("p2"), Type: gno.X("string")},
			{Name: gno.N("p3"), Type: gno.X("int64")},
		},
		[]gno.FieldTypeExpr{},
		true,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint8
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  string
				rp1 = reflect.ValueOf(&p1).Elem()
				p2  string
				rp2 = reflect.ValueOf(&p2).Elem()
				p3  int64
				rp3 = reflect.ValueOf(&p3).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 2, "")).TV, rp2)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 3, "")).TV, rp3)

			libs_std.X_bankerRemoveCoin(
				m,
				p0, p1, p2, p3)
		},
	},
	{
		"std",
		"emit",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
			{Name: gno.N("p1"), Type: gno.X("[]string")},
		},
		[]gno.FieldTypeExpr{},
		true,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  []string
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			libs_std.X_emit(
				m,
				p0, p1)
		},
	},
	{
		"std",
		"AssertOriginCall",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{},
		true,
		func(m *gno.Machine) {
			libs_std.AssertOriginCall(
				m,
			)
		},
	},
	{
		"std",
		"ChainID",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		true,
		func(m *gno.Machine) {
			r0 := libs_std.ChainID(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"ChainDomain",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		true,
		func(m *gno.Machine) {
			r0 := libs_std.ChainDomain(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"ChainHeight",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("int64")},
		},
		true,
		func(m *gno.Machine) {
			r0 := libs_std.ChainHeight(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"originSend",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("[]string")},
			{Name: gno.N("r1"), Type: gno.X("[]int64")},
		},
		true,
		func(m *gno.Machine) {
			r0, r1 := libs_std.X_originSend(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
		},
	},
	{
		"std",
		"originCaller",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		true,
		func(m *gno.Machine) {
			r0 := libs_std.X_originCaller(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"originPkgAddr",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		true,
		func(m *gno.Machine) {
			r0 := libs_std.X_originPkgAddr(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"callerAt",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("int")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		true,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  int
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_std.X_callerAt(
				m,
				p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"getRealm",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("int")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
			{Name: gno.N("r1"), Type: gno.X("string")},
		},
		true,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  int
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0, r1 := libs_std.X_getRealm(
				m,
				p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
		},
	},
	{
		"std",
		"assertCallerIsRealm",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{},
		true,
		func(m *gno.Machine) {
			libs_std.X_assertCallerIsRealm(
				m,
			)
		},
	},
	{
		"std",
		"setParamString",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
			{Name: gno.N("p1"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{},
		true,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  string
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			libs_std.X_setParamString(
				m,
				p0, p1)
		},
	},
	{
		"std",
		"setParamBool",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
			{Name: gno.N("p1"), Type: gno.X("bool")},
		},
		[]gno.FieldTypeExpr{},
		true,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  bool
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			libs_std.X_setParamBool(
				m,
				p0, p1)
		},
	},
	{
		"std",
		"setParamInt64",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
			{Name: gno.N("p1"), Type: gno.X("int64")},
		},
		[]gno.FieldTypeExpr{},
		true,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  int64
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			libs_std.X_setParamInt64(
				m,
				p0, p1)
		},
	},
	{
		"std",
		"setParamUint64",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
			{Name: gno.N("p1"), Type: gno.X("uint64")},
		},
		[]gno.FieldTypeExpr{},
		true,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  uint64
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			libs_std.X_setParamUint64(
				m,
				p0, p1)
		},
	},
	{
		"std",
		"setParamBytes",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
			{Name: gno.N("p1"), Type: gno.X("[]byte")},
		},
		[]gno.FieldTypeExpr{},
		true,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  []byte
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			libs_std.X_setParamBytes(
				m,
				p0, p1)
		},
	},
	{
		"testing",
		"unixNano",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("int64")},
		},
		false,
		func(m *gno.Machine) {
			r0 := libs_testing.X_unixNano()

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"testing",
		"matchString",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
			{Name: gno.N("p1"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("bool")},
			{Name: gno.N("r1"), Type: gno.X("error")},
		},
		false,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  string
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			r0, r1 := libs_testing.X_matchString(p0, p1)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
		},
	},
	{
		"time",
		"now",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("int64")},
			{Name: gno.N("r1"), Type: gno.X("int32")},
			{Name: gno.N("r2"), Type: gno.X("int64")},
		},
		true,
		func(m *gno.Machine) {
			r0, r1, r2 := libs_time.X_now(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r2).Elem(),
			))
		},
	},
	{
		"time",
		"loadFromEmbeddedTZData",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("[]byte")},
			{Name: gno.N("r1"), Type: gno.X("bool")},
		},
		false,
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0, r1 := libs_time.X_loadFromEmbeddedTZData(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
		},
	},
}

var initOrder = [...]string{
	"errors",
	"internal/bytealg",
	"io",
	"unicode",
	"unicode/utf8",
	"bytes",
	"strings",
	"bufio",
	"crypto/bech32",
	"encoding/binary",
	"math/bits",
	"math",
	"crypto/chacha20/chacha",
	"crypto/cipher",
	"crypto/chacha20",
	"strconv",
	"crypto/chacha20/rand",
	"crypto/ed25519",
	"crypto/sha256",
	"encoding",
	"encoding/base32",
	"encoding/base64",
	"encoding/csv",
	"encoding/hex",
	"hash",
	"hash/adler32",
	"html",
	"math/overflow",
	"math/rand",
	"path",
	"sort",
	"net/url",
	"regexp/syntax",
	"regexp",
	"std",
	"testing",
	"time",
	"unicode/utf16",
}

// InitOrder returns the initialization order of the standard libraries.
// This is calculated starting from the list of all standard libraries and
// iterating through each: if a package depends on an unitialized package, that
// is processed first, and so on recursively; matching the behaviour of Go's
// [program initialization].
//
// [program initialization]: https://go.dev/ref/spec#Program_initialization
func InitOrder() []string {
	return initOrder[:]
}
