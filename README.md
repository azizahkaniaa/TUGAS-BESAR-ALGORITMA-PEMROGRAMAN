# TUGAS-BESAR-ALGORITMA-PEMROGRAMAN
Sistem Pemesanan Makanan di Kantin 

# Deskripsi Program Manajemen Pesanan Makanan
Program ini adalah sistem manajemen pesanan makanan sederhana yang ditulis dalam bahasa pemrograman Go. Program ini dirancang untuk membantu penjual  dalam mengelola daftar pesanan, menambahkan item ke pesanan, melihat struk, serta melakukan analisis dasar seperti pencarian dan pengurutan data pesanan. Data disimpan dalam struktur array statis, yang berarti kapasitas maksimum pesanan dan item per pesanan telah ditentukan di awal.

# Fitur-fitur Utama
1. Pengelolaan Pesanan (kelolaPesananMenu())
•	Input ID Pesanan Baru (inputIDPesanan()): Memungkinkan pengguna untuk membuat pesanan baru dengan menetapkan ID unik. Program akan memeriksa apakah ID sudah ada dan memberikan pesan jika kapasitas pesanan penuh.
•	Tambah Item ke Pesanan (pilihMakanan()): Pengguna dapat memilih ID pesanan yang sudah ada, lalu memilih item makanan dari daftar menu yang tersedia (Nasi Goreng, Mie Ayam, Bakso, Soto Ayam, Es Teh, Es Jeruk) beserta jumlahnya. Fitur ini juga memvalidasi apakah pesanan sudah mencapai batas maksimum item.
•	Hapus Data Pesanan (hapusPesanan()): Memungkinkan pengguna untuk menghapus seluruh pesanan berdasarkan ID yang dimasukkan.
2. Melihat Data Pesanan (lihatDataPesananMenu())
•	Cetak Semua Struk Belanja (cetakStruk()): Menampilkan semua pesanan yang telah dibuat dalam format struk. Setiap struk mencakup ID pesanan, daftar item, subtotal untuk setiap item, dan total harga pesanan. Fungsi ini juga mendemonstrasikan penghitungan total harga menggunakan rekursi.
•	Cari ID Pesanan (Sequential Search) (cariIDPesananSequential()): Memungkinkan pengguna untuk mencari detail pesanan spesifik dengan memasukkan ID pesanan. Pencarian ini dilakukan secara sequential search (pencarian berurutan).
•	Cari Pesanan berdasarkan Total Harga (Binary Search) (cariTotalHargaPesananBinary()): Pengguna dapat mencari pesanan berdasarkan total harganya. Sebelum melakukan pencarian, data pesanan akan diurutkan berdasarkan total harga menggunakan insertion sort untuk memungkinkan pencarian binary search yang efisien.
3. Pengurutan Data Pesanan (urutkanDataPesananMenu())
•	Urutkan Harga Termahal (Selection Sort) (urutkanHargaTermahal()): Mengurutkan dan menampilkan daftar pesanan berdasarkan total harga dari yang tertinggi ke terendah menggunakan algoritma selection sort.
•	Urutkan Harga Termurah (Insertion Sort) (urutkanHargaTermurah()): Mengurutkan dan menampilkan daftar pesanan berdasarkan total harga dari yang terendah ke tertinggi menggunakan algoritma insertion sort.
4. Analisis Data Pesanan (analisisDataPesananMenu())
•	Cari Nilai Ekstrim (Pesanan Termahal & Termurah) (findExtremeOrderValues()): Menganalisis semua pesanan untuk menemukan dan menampilkan ID pesanan dengan total harga paling mahal dan paling murah. Fitur ini juga menyertakan demonstrasi penghitungan total harga pesanan untuk setiap entri menggunakan fungsi rekursif.

# Cara Kerja Umum
1.	Struktur Data: Program menggunakan struct untuk mendefinisikan MenuMakanan, ItemMakanan, dan Pesanan. Pesanan dan menu disimpan dalam array statis global.
2.	Input Pengguna: Fungsi bacaInt() yang telah ditingkatkan digunakan untuk membaca input integer dari pengguna dengan penanganan kesalahan yang lebih baik, memastikan hanya angka yang valid yang diterima.
3.	Tampilan Berbingkai: Fungsi cetakDenganBingkai() digunakan secara ekstensif untuk menyajikan semua output menu dan pesan kepada pengguna dalam format bingkai teks yang rapi dan mudah dibaca. Ini meningkatkan pengalaman pengguna secara keseluruhan.
4.	Algoritma Pencarian & Pengurutan: Program ini secara eksplisit mengimplementasikan berbagai algoritma dasar seperti Sequential Search, Binary Search, Selection Sort, dan Insertion Sort untuk mengelola dan menganalisis data pesanan.
5.	Rekursi: Fungsi totalHargaRecursive() mendemonstrasikan konsep rekursi untuk menghitung total harga suatu pesanan.

