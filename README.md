# Golang Interface

Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan menggunakan nama tertentu.

Interface merupakan sebuah tipe data, zero value dari sebuah interface adalah `nil`.

## membuat sebuah interface

untuk menerapkan sebuah interface, yang perlu dilakukan terlebih dahulu adalah membuatnya, keyword `type` dan `interface` digunakan untuk membuat sebuah interface.

<details>
  <summary>contoh membuat interface</summary>

```go
type hitung interface {
  luas() float64
  keliling() float64
}
```
</details><br>

pada conoth kode di atas, interface `hitung` memiliki 2 buah method, `lus()` dan `keliling()`. 
interface ini nantinya digunakan sebagai tipe data dari sebuah variabel, dimana variabel tersebut akan berisi data struct yang akan dibuat.
Dimana struct yang akan dibuat berisi method minimal yang tercantum saat pendeklarasian interface

<details>
  <summary>mempersiapkan struct dan method bangun datar</summary>

```go

// mempersiapkan struct dan method untuk bangun datar lingkaran
type lingkaran struct {
  diameter float64
}

func (l lingkaran) jarijari() float64 {
  return l.diameter / 2
}

func (l lingkaran) luas() float64 {
  return math.Pi * math.Pow(l.jarijari(), 2)
}

func (l lingkaran) keliling() float64 {
  return math.Pi * l.diameter
}

// mempersiapkan struct dan method untuk bangun datar persegi
type persegi struct {
  sisi float64
}

func (p persegi) luas() float64 {
  return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
  return p.sisi * 4
}
```
</details><br>

perbedaan antara struct persegi dan lingkaran di atas terletak pada method `jarijari()`, struct persegi tidak memiliki method tersebut, namun variabel hasil dari 2 struct ini akan tetap bisa ditampung oleh variabel yang bertipe data `interface hitung`.

<details>
  <summary>implementasi perhitungan di <code>func main</code></summary>

```go
func main() {
  var bangunDatar hitung

  bangunDatar = persegi{sisi: 10.0}
  fmt.Println("=====> persegi")
  fmt.Println("luas : ", bangunDatar.luas())
  fmt.Println("keliling : ", bangunDatar.keliling())

  bangunDatar = lingkaran{diameter: 14.0}
  fmt.Println("=====> lingkaran")
  fmt.Println("luas : ", bangunDatar.luas())
  fmt.Println("keliling : ", bangunDatar.keliling())
  fmt.Println("jari jari : ", bangunDatar.(lingkaran).jarijari())
}
```
</details>