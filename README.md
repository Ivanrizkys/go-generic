# Generic

Fitur generic di golang memungkinkan kita untuk membuat gabungan tipe data tertentu agar tipe data yang kita buat menjadi lebih dinamis. Hal ini sangat berguna ketika kita membuat function yang menerima parameter yang tipe datanya tidak fix atau bisa support akan beberapa tipe data sekaligus. Sebenarnya kita bisa menggunakan interface kosong untuk mengakalinya, namun kita harus memparsingnya lagi di lain waktu dan bisa saja proses parsing itu gagal. Oleh karena itu, kita bisa mengatasi masalah tersebut dengan memanfaatkan fitur golang generic.

## Tipe Data Any

Tipe data any di golang ini merupakan representasi dari interface{} yang mana artinya dapat support semua tipe data.

## Tipe Data Comparable

Tipe data comparable merupakan representasi tipe data yang dapat di bandingkan (booelan, numbers, strings, pointers, channels array dari tipe data comparable dan struct dari tipe data comparable).

## Generic Function

Cara membuat generic function di golang ini cukup mudah, kita hanya perlu merepresentasikan variable generic di dalam tanda [] setelah nama function. Biasanya penamaan variabel menggunakan satu huruf besar (T, K, V, dll) diikuti denga type constraint nya.

```go
// singel type parameter
func Length[T any](param T) T {
	return param
}

func TestSample(t *testing.T) {
    // pemanggilan generic function
	str := Length[string]("yahaha")
	number := Length[int](30)

	log.Println(str)
	log.Println(number)
}

// multiple type parameter
func MultipleTypeParameter[T any, K any](param1 T, param2 K) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleType(t *testing.T) {
	MultipleTypeParameter[string, int]("yagura", 9)
}

```

### Type Comparable

Jika kita ingin membandingkan type data dari generic function, maka kita tidak bisa merepresetasikan type constraint dengan menggunakan any karena tipe data any ini memang tidak bisa dibandingkan. Oleh karena itu, sebagai gantinya kita bisa menggunakan type data comparable.

```go
func IsSame[T comparable](value1, value2 T) bool {
    // jika T merupakan tipe data any maka kode dibawah ini akan error
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	log.Println("apakah ini sama", IsSame[string]("yahaha", "yahaha"))
	log.Println("apakah ini sama", IsSame[int](7, 9))
}
```

### Type Inheritance

Kita juga bisa menggunakan interface sebagai type constraint  dari generic function yang akan kita buat. Teknik ini desebut dengan type inheritance

```go
type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

// * Manager mengimplementasikan interface Employee
type Manager interface {
	GetName() string
	GetManagerName() string
}

// * MyManager mengimplementasikan interface Employee
type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

// * VicePrecident mengimplementasikan interface Employee
type VicePrecident interface {
	GetName() string
	GetVicePrecidentName() string
}

// * MyVicePrecident mengimplementasikan interface Employee
type MyVicePrecident struct {
	Name string
}

func (m *MyVicePrecident) GetName() string {
	return m.Name
}

func (m *MyVicePrecident) GetVicePrecidentName() string {
	return m.Name
}

func TestTypeInheritance(t *testing.T) {
	assert.Equal(t, "Ivan", GetName[Manager](&MyManager{Name: "Ivan"}))
	assert.Equal(t, "Aisyah", GetName[VicePrecident](&MyVicePrecident{Name: "Aisyah"}))
}
```

### Type Set

Dengan type set kita bisa membuat kumpulan type constraint yang support dengan tipe data tertentu yang kita buat sendiri. Jika type data yang kita deklarasikan semuanya support dengan perintah matematika (int, float) maka perintah itu juga bisa kita gunakan di function generic yang kita buat

```go
func Min[T Number](first, second T) T {
    // jika interface number terdapat tipe misalkan bool
    // maka kode dibawah akan error
	if first < second {
		return first
	} else {
		return second
	}
}

func TestTypeSet(t *testing.T) {
	assert.Equal(t, float32(100), Min[float32](float32(300), float32(100)))
	assert.Equal(t, int32(50), Min[int32](int32(50), int32(600)))
}
```

## Generic Type

Di golang juga memungkinkan untuk membuat generic type dengan cara yang tidak jauh beda dengan saat membuat generic function

```go
type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		log.Println(value)
	}
}

func TestTypeTest(t *testing.T) {
	name := Bag[string]{"lahura", "yahihi", "kuyao", "siti"}
	PrintBag(name)
	assert.Equal(t, true, true)
}
```