# Golang Interface

Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan menggunakan nama tertentu.

Interface merupakan sebuah tipe data, sero value dari sebuah interface adalah `nil`.

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
interface ini nantinya digunakan sebagai tipe data dari sebuah variabel.