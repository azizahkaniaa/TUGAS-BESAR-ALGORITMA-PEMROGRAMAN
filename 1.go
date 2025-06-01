package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Konstanta untuk ukuran array statis
const MAX_PESANAN = 100
const MAX_ITEMS_PER_PESANAN = 10
const MAX_MENU = 10

// Untuk menyatakan perkiraan maksimal jumlah baris tampilan yang dibutuhkan untuk mencetak semua pesanan beserta item-itemnya
const MAX_FRAME_LINES = MAX_PESANAN * (MAX_ITEMS_PER_PESANAN + 5) 

// Struktur data
type MenuMakanan struct {
	Nomor int
	Nama  string
	Harga int
}

type ItemMakanan struct {
	Nama        string
	Jumlah      int
	HargaSatuan int
}

type Pesanan struct {
	ID          int
	Items       [MAX_ITEMS_PER_PESANAN]ItemMakanan
	JumlahItems int 
}

// Variabel global
var daftarPesanan [MAX_PESANAN]Pesanan
var jumlahPesananSaatIni int = 0

// Menyimpan daftar makanan beserta harga, dan menunjukkan jumlah menu yang saat ini tersedia
var menu = [MAX_MENU]MenuMakanan{
	{1, "Nasi Goreng", 15000},
	{2, "Mie Ayam", 12000},
	{3, "Bakso", 13000},
	{4, "Soto Ayam", 14000},
	{5, "Es Teh", 5000},
	{6, "Es Jeruk", 7000},
}
var jumlahMenuAktual int = 6

// Dipakai untuk menghitung jumlah huruf atau karakter dalam sebuah teks
func getStringRuneCount(s string) int {
	return len([]rune(s)) 
}

// Fungsi untuk mencetak teks dengan bingkai
func cetakDenganBingkai(teks [MAX_FRAME_LINES]string, jumlahBaris int) {
	var maxPanjang int = 0
	var i int
	for i = 0; i < jumlahBaris; i++ {
		var baris string = teks[i]
		var currentBarisPanjang int = getStringRuneCount(baris) 
		if currentBarisPanjang > maxPanjang {
			maxPanjang = currentBarisPanjang
		}
	}

	fmt.Print("┌")
	var k_border int
	for k_border = 0; k_border < maxPanjang+2; k_border++ {
		fmt.Print("─")
	}
	fmt.Println("┐")

	for i = 0; i < jumlahBaris; i++ {
		var baris string = teks[i]
		var barisPanjang int = getStringRuneCount(baris) 
		var padding int = maxPanjang - barisPanjang
		fmt.Print("│ ", baris)
		var k_padding int
		for k_padding = 0; k_padding < padding; k_padding++ {
			fmt.Print(" ")
		}
		fmt.Println(" │")
	}

	fmt.Print("└")
	for k_border = 0; k_border < maxPanjang+2; k_border++ {
		fmt.Print("─")
	}
	fmt.Println("┘")
}

// Fungsi untuk membaca input integer dari pengguna
func bacaInt() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Masukkan angka: ")
		inputString, _ := reader.ReadString('\n')
		inputString = strings.TrimSpace(inputString)
		input, err := strconv.Atoi(inputString)
		if err != nil {
			fmt.Println("Input tidak valid. Harap masukkan angka.")
		} else {
			return input
		}
	}
}

// Fungsi utama program
func main() {
	for {
		var menuUtamaTeks [MAX_FRAME_LINES]string 
		var jumlahMenuUtama int = 0

		menuUtamaTeks[jumlahMenuUtama] = "Created by Alya dan Azizah"
		jumlahMenuUtama++
		menuUtamaTeks[jumlahMenuUtama] = "======= MENU UTAMA ======="
		jumlahMenuUtama++
		menuUtamaTeks[jumlahMenuUtama] = "1. Kelola Pesanan"
		jumlahMenuUtama++
		menuUtamaTeks[jumlahMenuUtama] = "2. Lihat Data Pesanan"
		jumlahMenuUtama++
		menuUtamaTeks[jumlahMenuUtama] = "3. Urutkan Data Pesanan"
		jumlahMenuUtama++
		menuUtamaTeks[jumlahMenuUtama] = "4. Analisis Data Pesanan"
		jumlahMenuUtama++
		menuUtamaTeks[jumlahMenuUtama] = "5. Keluar"
		jumlahMenuUtama++

		cetakDenganBingkai(menuUtamaTeks, jumlahMenuUtama)
		fmt.Print("Pilih menu (1-5): ")
		var pilihan int = bacaInt()
		fmt.Println()

		switch pilihan {
		case 1:
			kelolaPesananMenu()
		case 2:
			lihatDataPesananMenu()
		case 3:
			urutkanDataPesananMenu()
		case 4:
			analisisDataPesananMenu()
		case 5:
			var exitMessageTeks [MAX_FRAME_LINES]string
			var exitMessageCount int = 0
			exitMessageTeks[exitMessageCount] = "Terima kasih! Program selesai."
			exitMessageCount++
			cetakDenganBingkai(exitMessageTeks, exitMessageCount)
			return
		default:
			var invalidChoiceMessageTeks [MAX_FRAME_LINES]string
			var invalidChoiceMessageCount int = 0
			invalidChoiceMessageTeks[invalidChoiceMessageCount] = "Pilihan tidak valid."
			invalidChoiceMessageCount++
			cetakDenganBingkai(invalidChoiceMessageTeks, invalidChoiceMessageCount)
		}
	}
}

// Untuk menampilkan pilihan pengelolaan pesanan dan menjalankan aksi sesuai pilihan pengguna, seperti menambah atau menghapus pesanan
func kelolaPesananMenu() { 
	for {
		var menuKelolaTeks [MAX_FRAME_LINES]string
		var jumlahMenuKelola int = 0
		menuKelolaTeks[jumlahMenuKelola] = "--- Kelola Pesanan ---"
		jumlahMenuKelola++
		menuKelolaTeks[jumlahMenuKelola] = "1. Input ID Pesanan Baru"
		jumlahMenuKelola++
		menuKelolaTeks[jumlahMenuKelola] = "2. Tambah Item ke Pesanan"
		jumlahMenuKelola++
		menuKelolaTeks[jumlahMenuKelola] = "3. Hapus Data Pesanan"
		jumlahMenuKelola++
		menuKelolaTeks[jumlahMenuKelola] = "4. Kembali ke Menu Utama"
		jumlahMenuKelola++

		cetakDenganBingkai(menuKelolaTeks, jumlahMenuKelola)
		fmt.Print("Pilih sub-menu (1-4): ")
		var pilihan int = bacaInt()
		fmt.Println()

		switch pilihan {
		case 1:
			inputIDPesanan()
		case 2:
			pilihMakanan()
		case 3:
			hapusPesanan()
		case 4:
			return
		default:
			var invalidChoiceMessageTeks [MAX_FRAME_LINES]string
			var invalidChoiceMessageCount int = 0
			invalidChoiceMessageTeks[invalidChoiceMessageCount] = "Pilihan tidak valid."
			invalidChoiceMessageCount++
			cetakDenganBingkai(invalidChoiceMessageTeks, invalidChoiceMessageCount)
		}
	}
}

// Digunakan untuk menampilkan menu pilihan melihat data pesanan, seperti mencetak struk, mencari pesanan 
// berdasarkan ID atau total harga, dan kembali ke menu utama.
func lihatDataPesananMenu() {
	var messageTeksIfEmpty [MAX_FRAME_LINES]string
	var messageCountIfEmpty int = 0
	if jumlahPesananSaatIni == 0 {
		messageTeksIfEmpty[messageCountIfEmpty] = "Belum ada pesanan untuk dilihat."
		messageCountIfEmpty++
		cetakDenganBingkai(messageTeksIfEmpty, messageCountIfEmpty)
		return
	}
	for {
		var menuLihatTeks [MAX_FRAME_LINES]string
		var jumlahMenuLihat int = 0
		menuLihatTeks[jumlahMenuLihat] = "--- Lihat Data Pesanan ---"
		jumlahMenuLihat++
		menuLihatTeks[jumlahMenuLihat] = "1. Cetak Semua Struk Belanja"
		jumlahMenuLihat++
		menuLihatTeks[jumlahMenuLihat] = "2. Cari ID Pesanan (Sequential Search)"
		jumlahMenuLihat++
		menuLihatTeks[jumlahMenuLihat] = "3. Cari Pesanan berdasarkan Total Harga (Binary Search)"
		jumlahMenuLihat++
		menuLihatTeks[jumlahMenuLihat] = "4. Kembali ke Menu Utama"
		jumlahMenuLihat++

		cetakDenganBingkai(menuLihatTeks, jumlahMenuLihat)
		fmt.Print("Pilih sub-menu (1-4): ")
		var pilihan int = bacaInt()
		fmt.Println()

		switch pilihan {
		case 1:
			cetakStruk()
		case 2:
			cariIDPesananSequential()
		case 3:
			cariTotalHargaPesananBinary()
		case 4:
			return
		default:
			var invalidChoiceMessageTeks [MAX_FRAME_LINES]string
			var invalidChoiceMessageCount int = 0
			invalidChoiceMessageTeks[invalidChoiceMessageCount] = "Pilihan tidak valid."
			invalidChoiceMessageCount++
			cetakDenganBingkai(invalidChoiceMessageTeks, invalidChoiceMessageCount)
		}
	}
}

// Menampilkan menu untuk mengurutkan data pesanan berdasarkan harga termahal menggunakan selection sort atau 
// harga termurah menggunakan insertion sort, serta menyediakan opsi kembali ke menu utama
func urutkanDataPesananMenu() {
	var messageTeksIfEmpty [MAX_FRAME_LINES]string
	var messageCountIfEmpty int = 0
	if jumlahPesananSaatIni == 0 {
		messageTeksIfEmpty[messageCountIfEmpty] = "Belum ada pesanan untuk diurutkan."
		messageCountIfEmpty++
		cetakDenganBingkai(messageTeksIfEmpty, messageCountIfEmpty)
		return
	}
	for {
		var menuUrutkanTeks [MAX_FRAME_LINES]string
		var jumlahMenuUrutkan int = 0
		menuUrutkanTeks[jumlahMenuUrutkan] = "--- Urutkan Data Pesanan ---"
		jumlahMenuUrutkan++
		menuUrutkanTeks[jumlahMenuUrutkan] = "1. Urutkan Harga Termahal (Selection Sort)"
		jumlahMenuUrutkan++
		menuUrutkanTelahDitemukan := false
		for i := 0; i < jumlahMenuUrutkan; i++ {
			if menuUrutkanTeks[i] == "2. Urutkan Harga Termurah (Insertion Sort)" {
				menuUrutkanTelahDitemukan = true
				break
			}
		}
		if !menuUrutkanTelahDitemukan {
			menuUrutkanTeks[jumlahMenuUrutkan] = "2. Urutkan Harga Termurah (Insertion Sort)"
			jumlahMenuUrutkan++
		}
		menuUrutkanTeks[jumlahMenuUrutkan] = "3. Kembali ke Menu Utama"
		jumlahMenuUrutkan++

		cetakDenganBingkai(menuUrutkanTeks, jumlahMenuUrutkan)
		fmt.Print("Pilih sub-menu (1-3): ")
		var pilihan int = bacaInt()
		fmt.Println()

		switch pilihan {
		case 1:
			urutkanHargaTermahal()
		case 2:
			urutkanHargaTermurah()
		case 3:
			return
		default:
			var invalidChoiceMessageTeks [MAX_FRAME_LINES]string
			var invalidChoiceMessageCount int = 0
			invalidChoiceMessageTeks[invalidChoiceMessageCount] = "Pilihan tidak valid."
			invalidChoiceMessageCount++
			cetakDenganBingkai(invalidChoiceMessageTeks, invalidChoiceMessageCount)
		}
	}
}

// Menampilkan menu analisis untuk mencari nilai ekstrem dari data pesanan, 
// yaitu pesanan dengan harga termahal dan termurah, atau kembali ke menu utama
func analisisDataPesananMenu() {
	var messageTeksIfEmpty [MAX_FRAME_LINES]string
	var messageCountIfEmpty int = 0
	if jumlahPesananSaatIni == 0 {
		messageTeksIfEmpty[messageCountIfEmpty] = "Belum ada pesanan untuk dianalisis."
		messageCountIfEmpty++
		cetakDenganBingkai(messageTeksIfEmpty, messageCountIfEmpty)
		return
	}
	for {
		var menuAnalisisTeks [MAX_FRAME_LINES]string
		var jumlahMenuAnalisis int = 0
		menuAnalisisTeks[jumlahMenuAnalisis] = "--- Analisis Data Pesanan ---"
		jumlahMenuAnalisis++
		menuAnalisisTeks[jumlahMenuAnalisis] = "1. Cari Nilai Ekstrim (Pesanan Termahal & Termurah)"
		jumlahMenuAnalisis++
		menuAnalisisTeks[jumlahMenuAnalisis] = "2. Kembali ke Menu Utama"
		jumlahMenuAnalisis++

		cetakDenganBingkai(menuAnalisisTeks, jumlahMenuAnalisis)
		fmt.Print("Pilih sub-menu (1-2): ")
		var pilihan int = bacaInt()
		fmt.Println()

		switch pilihan {
		case 1:
			findExtremeOrderValues()
		case 2:
			return
		default:
			var invalidChoiceMessageTeks [MAX_FRAME_LINES]string
			var invalidChoiceMessageCount int = 0
			invalidChoiceMessageTeks[invalidChoiceMessageCount] = "Pilihan tidak valid."
			invalidChoiceMessageCount++
			cetakDenganBingkai(invalidChoiceMessageTeks, invalidChoiceMessageCount)
		}
	}
}

// Digunakan untuk menambahkan ID pesanan baru selama kapasitas 
// belum penuh dan ID tersebut belum digunakan sebelumnya
func inputIDPesanan() {
	var messageTeks [MAX_FRAME_LINES]string
	var messageCount int = 0

	if jumlahPesananSaatIni >= MAX_PESANAN {
		messageTeks[messageCount] = "Kapasitas pesanan penuh."
		messageCount++
		cetakDenganBingkai(messageTeks, messageCount)
		return
	}

	fmt.Print("Masukkan ID Pesanan: ")
	var id int = bacaInt()

	var idSudahAda bool = false
	var i int
	for i = 0; i < jumlahPesananSaatIni; i++ {
		if daftarPesanan[i].ID == id {
			idSudahAda = true
			break
		}
	}

	if idSudahAda {
		messageTeks[messageCount] = "ID Pesanan sudah ada. Silakan gunakan ID lain."
	} else {
		daftarPesanan[jumlahPesananSaatIni] = Pesanan{ID: id, JumlahItems: 0}
		jumlahPesananSaatIni++
		messageTeks[messageCount] = "ID disimpan. Sekarang pilih makanan di menu Kelola Pesanan -> Tambah Item."
	}
	cetakDenganBingkai(messageTeks, messageCount)
}

// Digunakan untuk menambahkan item makanan ke pesanan 
// tertentu berdasarkan ID yang sudah dimasukkan sebelumnya
func pilihMakanan() {
	var messageTeks [MAX_FRAME_LINES]string
	var messageCount int = 0

	if jumlahPesananSaatIni == 0 {
		messageTeks[messageCount] = "Harap masukkan ID Pesanan terlebih dahulu (Kelola Pesanan -> Input ID)."
		messageCount++
		cetakDenganBingkai(messageTeks, messageCount)
		return
	}

	var pesananTargetIndex int = -1
	fmt.Print("Masukkan ID Pesanan yang akan ditambahkan item: ")
	var idTarget int = bacaInt()

	var foundTargetPesanan bool = false
	var i int
	for i = 0; i < jumlahPesananSaatIni; i++ {
		if daftarPesanan[i].ID == idTarget {
			pesananTargetIndex = i
			foundTargetPesanan = true
			break
		}
	}

	if !foundTargetPesanan {
		messageTeks[messageCount] = fmt.Sprintf("ID Pesanan %d tidak ditemukan.", idTarget)
		messageCount++
		cetakDenganBingkai(messageTeks, messageCount)
		return
	}

	var pesananSaatIni *Pesanan = &daftarPesanan[pesananTargetIndex]

	if pesananSaatIni.JumlahItems >= MAX_ITEMS_PER_PESANAN {
		messageTeks[messageCount] = "Jumlah item maksimum untuk pesanan ini telah tercapai."
		messageCount++
		cetakDenganBingkai(messageTeks, messageCount)
		return
	}

	var teksMenuDisplay [MAX_FRAME_LINES]string 
	var idxTeksMenu int = 0

	teksMenuDisplay[idxTeksMenu] = "Pilih Makanan:"
	idxTeksMenu++

	var k int
	for k = 0; k < jumlahMenuAktual; k++ {
		var item MenuMakanan = menu[k]
		if item.Nomor != 0 { 
			if idxTeksMenu < MAX_FRAME_LINES { 
				teksMenuDisplay[idxTeksMenu] = fmt.Sprintf("%d. %s - Rp%d", item.Nomor, item.Nama, item.Harga)
				idxTeksMenu++
			}
		}
	}
	cetakDenganBingkai(teksMenuDisplay, idxTeksMenu)

	fmt.Print("Masukkan nomor makanan: ")
	var pilihanMenu int = bacaInt()

	var itemDipilih *MenuMakanan
	var foundItemMenu bool = false
	var j int
	for j = 0; j < jumlahMenuAktual; j++ {
		if menu[j].Nomor == pilihanMenu {
			itemDipilih = &menu[j]
			foundItemMenu = true
			break
		}
	}

	if !foundItemMenu || itemDipilih == nil || itemDipilih.Nomor == 0 {
		messageTeks[messageCount] = "Menu tidak tersedia."
		messageCount++
		cetakDenganBingkai(messageTeks, messageCount)
		return
	}

	fmt.Print("Masukkan jumlah: ")
	var jumlahItem int = bacaInt()
	if jumlahItem <= 0 {
		messageTeks[messageCount] = "Jumlah item tidak valid."
		messageCount++
		cetakDenganBingkai(messageTeks, messageCount)
		return
	}

	pesananSaatIni.Items[pesananSaatIni.JumlahItems] = ItemMakanan{
		Nama:        itemDipilih.Nama,
		Jumlah:      jumlahItem,
		HargaSatuan: itemDipilih.Harga,
	}
	pesananSaatIni.JumlahItems++

	messageTeks[messageCount] = "Makanan berhasil ditambahkan ke pesanan."
	messageCount++
	cetakDenganBingkai(messageTeks, messageCount)
}

// digunakan untuk menampilkan struk belanja dari semua pesanan yang ada, 
//termasuk detail item, subtotal, dan total harga pesanan
func cetakStruk() {
	var strukTeks [MAX_FRAME_LINES]string
	var strukIdx int = 0

	if jumlahPesananSaatIni == 0 {
		strukTeks[strukIdx] = "Belum ada pesanan."
		strukIdx++
		cetakDenganBingkai(strukTeks, strukIdx)
		return
	}

	strukTeks[strukIdx] = "Struk Belanja:"
	strukIdx++

	var i int
	for i = 0; i < jumlahPesananSaatIni; i++ {
		if strukIdx >= MAX_FRAME_LINES-(MAX_ITEMS_PER_PESANAN+3) {
			break
		}
		var p Pesanan = daftarPesanan[i]
		strukTeks[strukIdx] = fmt.Sprintf("ID Pesanan: %d", p.ID)
		strukIdx++
		var totalPesanan int = 0
		if p.JumlahItems == 0 {
			if strukIdx < MAX_FRAME_LINES {
				strukTeks[strukIdx] = "  (Belum ada item makanan)"
				strukIdx++
			}
		} else {
			var j_item int
			for j_item = 0; j_item < p.JumlahItems; j_item++ {
				if strukIdx >= MAX_FRAME_LINES {
					break
				}
				var item ItemMakanan = p.Items[j_item]
				var subtotal int = item.Jumlah * item.HargaSatuan
				strukTeks[strukIdx] = fmt.Sprintf("  - %s x%d @ Rp%d = Rp%d", item.Nama, item.Jumlah, item.HargaSatuan, subtotal)
				strukIdx++
				totalPesanan += subtotal
			}
		}
		if strukIdx < MAX_FRAME_LINES {
			var totalPesananRecursive int = totalHargaRecursive(p, 0)
			strukTeks[strukIdx] = fmt.Sprintf("  Total untuk ID %d: Rp%d (Rekursif: Rp%d)", p.ID, totalPesanan, totalPesananRecursive)
			strukIdx++
		}
		if strukIdx < MAX_FRAME_LINES {
			strukTeks[strukIdx] = ""
			strukIdx++
		}
	}
	cetakDenganBingkai(strukTeks, strukIdx)
}

// Digunakan untuk menghitung total harga dari seluruh item 
// makanan dalam suatu pesanan berdasarkan jumlah dan harga satuannya
func hitungTotalHargaPesanan(items [MAX_ITEMS_PER_PESANAN]ItemMakanan, jumlahItemAktual int) int {
	var total int = 0
	var i int
	for i = 0; i < jumlahItemAktual; i++ {
		total += items[i].Jumlah * items[i].HargaSatuan
	}
	return total
}

// Menghitung total harga pesanan dengan menjumlahkan subtotal setiap 
// item mulai dari indeks saat ini hingga seluruh item habis diperiksa
func totalHargaRecursive(pesanan Pesanan, currentIndex int) int {
	if currentIndex >= pesanan.JumlahItems {
		return 0
	}
	var currentItem ItemMakanan = pesanan.Items[currentIndex]
	var subtotal int = currentItem.Jumlah * currentItem.HargaSatuan
	return subtotal + totalHargaRecursive(pesanan, currentIndex+1)
}

// Mengurutkan data pesanan berdasarkan total harga dari yang 
// paling mahal ke yang paling murah menggunakan algoritma selection sort
func urutkanHargaTermahal() {
	var messageTeks [MAX_FRAME_LINES]string
	var messageCount int = 0
	if jumlahPesananSaatIni == 0 {
		messageTeks[messageCount] = "Belum ada pesanan untuk diurutkan."
		messageCount++
		cetakDenganBingkai(messageTeks, messageCount)
		return
	}

	var dataUrut [MAX_PESANAN]Pesanan 
	var i_copy int
	for i_copy = 0; i_copy < jumlahPesananSaatIni; i_copy++ {
		dataUrut[i_copy] = daftarPesanan[i_copy]
	}

	var i int
	for i = 0; i < jumlahPesananSaatIni-1; i++ {
		var maxIdx int = i
		var totalMaxIdx int = hitungTotalHargaPesanan(dataUrut[maxIdx].Items, dataUrut[maxIdx].JumlahItems)
		var j int
		for j = i + 1; j < jumlahPesananSaatIni; j++ {
			var totalJ int = hitungTotalHargaPesanan(dataUrut[j].Items, dataUrut[j].JumlahItems)
			if totalJ > totalMaxIdx {
				maxIdx = j
				totalMaxIdx = totalJ
			}
		}
		if maxIdx != i {
			temp := dataUrut[i]
			dataUrut[i] = dataUrut[maxIdx]
			dataUrut[maxIdx] = temp
		}
	}

	var hasilUrutTeks [MAX_FRAME_LINES]string
	var hasilUrutIdx int = 0
	hasilUrutTeks[hasilUrutIdx] = "ID Pesanan urut dari total harga tertinggi:"
	hasilUrutIdx++

	var i_print int
	for i_print = 0; i_print < jumlahPesananSaatIni; i_print++ {
		if hasilUrutIdx >= MAX_FRAME_LINES {
			break
		}
		var p Pesanan = dataUrut[i_print]
		var total int = hitungTotalHargaPesanan(p.Items, p.JumlahItems)
		hasilUrutTeks[hasilUrutIdx] = fmt.Sprintf("ID: %d - Total: Rp%d", p.ID, total)
		hasilUrutIdx++
	}
	cetakDenganBingkai(hasilUrutTeks, hasilUrutIdx)
}

// Digunakan untuk menghapus pesanan berdasarkan ID yang 
// dimasukkan oleh pengguna jika pesanan tersebut ditemukan dalam daftar
func hapusPesanan() {
	var messageTeks [MAX_FRAME_LINES]string
	var messageCount int = 0

	if jumlahPesananSaatIni == 0 {
		messageTeks[messageCount] = "Belum ada pesanan untuk dihapus."
		messageCount++
		cetakDenganBingkai(messageTeks, messageCount)
		return
	}
	fmt.Print("Masukkan ID pesanan yang ingin dihapus: ")
	var idHapus int = bacaInt()

	var indexDitemukan int = -1
	var foundPesananHapus bool = false
	var i int
	for i = 0; i < jumlahPesananSaatIni; i++ {
		if daftarPesanan[i].ID == idHapus {
			indexDitemukan = i
			foundPesananHapus = true
			break
		}
	}

	if foundPesananHapus {
		var k int
		for k = indexDitemukan; k < jumlahPesananSaatIni-1; k++ {
			daftarPesanan[k] = daftarPesanan[k+1]
		}
		daftarPesanan[jumlahPesananSaatIni-1] = Pesanan{}
		jumlahPesananSaatIni--
		messageTeks[messageCount] = "Pesanan berhasil dihapus."
		messageCount++
	} else {
		messageTeks[messageCount] = fmt.Sprintf("ID Pesanan %d tidak ditemukan.", idHapus)
		messageCount++
	}
	cetakDenganBingkai(messageTeks, messageCount)
}

// Digunakan untuk mencari dan menampilkan pesanan dengan total 
// harga paling mahal dan paling murah menggunakan pendekatan rekursif
func findExtremeOrderValues() {
	var messageTeks [MAX_FRAME_LINES]string
	var messageCount int = 0

	if jumlahPesananSaatIni == 0 {
		messageTeks[messageCount] = "Belum ada pesanan untuk mencari nilai ekstrem."
		messageCount++
		cetakDenganBingkai(messageTeks, messageCount)
		return
	}

	minTotal := totalHargaRecursive(daftarPesanan[0], 0)
	maxTotal := totalHargaRecursive(daftarPesanan[0], 0)
	minID := daftarPesanan[0].ID
	maxID := daftarPesanan[0].ID

	var i_loop int
	for i_loop = 1; i_loop < jumlahPesananSaatIni; i_loop++ {
		currentTotal := totalHargaRecursive(daftarPesanan[i_loop], 0)
		if currentTotal < minTotal {
			minTotal = currentTotal
			minID = daftarPesanan[i_loop].ID
		}
		if currentTotal > maxTotal {
			maxTotal = currentTotal
			maxID = daftarPesanan[i_loop].ID
		}
	}

	var outputTeks [MAX_FRAME_LINES]string
	var outputIdx int = 0

	outputTeks[outputIdx] = "Analisis Nilai Ekstrim:"
	outputIdx++
	if outputIdx < MAX_FRAME_LINES {
		outputTeks[outputIdx] = fmt.Sprintf("Pesanan Termahal: ID %d dengan total Rp%d", maxID, maxTotal)
		outputIdx++
	}
	if outputIdx < MAX_FRAME_LINES {
		outputTeks[outputIdx] = fmt.Sprintf("Pesanan Termurah: ID %d dengan total Rp%d", minID, minTotal)
		outputIdx++
	}
	if outputIdx < MAX_FRAME_LINES {
		outputTeks[outputIdx] = ""
		outputIdx++
	}
	if outputIdx < MAX_FRAME_LINES {
		outputTeks[outputIdx] = "Demonstrasi Rekursi (Detail untuk setiap pesanan):"
		outputIdx++
	}

	var i_print int
	for i_print = 0; i_print < jumlahPesananSaatIni; i_print++ {
		if outputIdx >= MAX_FRAME_LINES {
			break
		}
		var p Pesanan = daftarPesanan[i_print]
		var totalRecursive int = totalHargaRecursive(p, 0)
		outputTeks[outputIdx] = fmt.Sprintf("ID %d Total Rekursif: Rp%d", p.ID, totalRecursive)
		outputIdx++
	}
	cetakDenganBingkai(outputTeks, outputIdx)
}

// Digunakan untuk mengurutkan dan menampilkan daftar pesanan berdasarkan total harga 
// dari yang paling murah ke yang paling mahal menggunakan algoritma insertion sort
func urutkanHargaTermurah() {
	var messageTeks [MAX_FRAME_LINES]string
	var messageCount int = 0
	if jumlahPesananSaatIni == 0 {
		messageTeks[messageCount] = "Belum ada pesanan untuk diurutkan."
		messageCount++
		cetakDenganBingkai(messageTeks, messageCount)
		return
	}

	var dataUrut [MAX_PESANAN]Pesanan
	var i_copy int
	for i_copy = 0; i_copy < jumlahPesananSaatIni; i_copy++ {
		dataUrut[i_copy] = daftarPesanan[i_copy]
	}

	var pass int = 1
	for pass < jumlahPesananSaatIni {
		var i int = pass
		var temp Pesanan = dataUrut[pass]
		var totalTemp int = hitungTotalHargaPesanan(temp.Items, temp.JumlahItems)
		for i > 0 {
			var totalPrev int = hitungTotalHargaPesanan(dataUrut[i-1].Items, dataUrut[i-1].JumlahItems)
			if totalTemp < totalPrev {
				dataUrut[i] = dataUrut[i-1]
				i--
			} else {
				break
			}
		}
		dataUrut[i] = temp
		pass++
	}

	var hasilUrutTeks [MAX_FRAME_LINES]string
	var hasilUrutIdx int = 0
	hasilUrutTeks[hasilUrutIdx] = "ID Pesanan urut dari total harga terendah:"
	hasilUrutIdx++

	var i_print int
	for i_print = 0; i_print < jumlahPesananSaatIni; i_print++ {
		if hasilUrutIdx >= MAX_FRAME_LINES {
			break
		}
		var p Pesanan = dataUrut[i_print]
		var total int = hitungTotalHargaPesanan(p.Items, p.JumlahItems)
		hasilUrutTeks[hasilUrutIdx] = fmt.Sprintf("ID: %d - Total: Rp%d", p.ID, total)
		hasilUrutIdx++
	}
	cetakDenganBingkai(hasilUrutTeks, hasilUrutIdx)
}

// Melakukan pencarian ID pesanan secara berurutan
// (sequential search) dan menampilkan detail pesanan jika ditemukan, atau pesan kesalahan jika tidak
func cariIDPesananSequential() {
	var messageTeks [MAX_FRAME_LINES]string 
	var messageCount int = 0

	if jumlahPesananSaatIni == 0 {
		messageTeks[messageCount] = "Belum ada pesanan untuk dicari."
		messageCount++
		cetakDenganBingkai(messageTeks, messageCount)
		return
	}

	fmt.Print("Masukkan ID Pesanan yang dicari: ")
	var idCari int = bacaInt()

	var found bool = false
	var foundIndex int = -1
	var i int
	for i = 0; i < jumlahPesananSaatIni; i++ {
		if daftarPesanan[i].ID == idCari {
			found = true
			foundIndex = i
			break
		}
	}

	if found {
		var foundMessageDisplayTeks [MAX_FRAME_LINES]string
		var foundMessageDisplayIdx int = 0

		if foundMessageDisplayIdx < MAX_FRAME_LINES {
			foundMessageDisplayTeks[foundMessageDisplayIdx] = fmt.Sprintf("ID Pesanan %d ditemukan pada indeks %d.", idCari, foundIndex)
			foundMessageDisplayIdx++
		}

		var p Pesanan = daftarPesanan[foundIndex]
		if foundMessageDisplayIdx < MAX_FRAME_LINES {
			foundMessageDisplayTeks[foundMessageDisplayIdx] = fmt.Sprintf("Detail Pesanan ID %d:", p.ID)
			foundMessageDisplayIdx++
		}

		var totalPesanan int = 0
		if p.JumlahItems == 0 {
			if foundMessageDisplayIdx < MAX_FRAME_LINES {
				foundMessageDisplayTeks[foundMessageDisplayIdx] = "  (Belum ada item makanan)"
				foundMessageDisplayIdx++
			}
		} else {
			var j_item int
			for j_item = 0; j_item < p.JumlahItems; j_item++ {
				if foundMessageDisplayIdx >= MAX_FRAME_LINES {
					break
				}
				var item ItemMakanan = p.Items[j_item]
				var subtotal int = item.Jumlah * item.HargaSatuan
				foundMessageDisplayTeks[foundMessageDisplayIdx] = fmt.Sprintf("  - %s x%d @ Rp%d = Rp%d", item.Nama, item.Jumlah, item.HargaSatuan, subtotal)
				foundMessageDisplayIdx++
				totalPesanan += subtotal
			}
		}
		if foundMessageDisplayIdx < MAX_FRAME_LINES {
			foundMessageDisplayTeks[foundMessageDisplayIdx] = fmt.Sprintf("  Total Pesanan: Rp%d", totalPesanan)
			foundMessageDisplayIdx++
		}
		cetakDenganBingkai(foundMessageDisplayTeks, foundMessageDisplayIdx)
	} else {
		messageTeks[messageCount] = fmt.Sprintf("ID Pesanan %d tidak ditemukan.", idCari)
		messageCount++
		cetakDenganBingkai(messageTeks, messageCount)
	}
}

// Melakukan pencarian total harga pesanan menggunakan algoritma binary search 
// setelah mengurutkan pesanan berdasarkan total harga secara ascending, dan menampilkan hasil pencarian
func cariTotalHargaPesananBinary() {
	var messageTeks [MAX_FRAME_LINES]string
	var messageCount int = 0

	if jumlahPesananSaatIni == 0 {
		messageTeks[messageCount] = "Belum ada pesanan untuk dicari."
		messageCount++
		cetakDenganBingkai(messageTeks, messageCount)
		return
	}

	var dataUrut [MAX_PESANAN]Pesanan
	var i_copy int
	for i_copy = 0; i_copy < jumlahPesananSaatIni; i_copy++ {
		dataUrut[i_copy] = daftarPesanan[i_copy]
	}

	var pass int = 1
	for pass < jumlahPesananSaatIni {
		var i int = pass
		var temp Pesanan = dataUrut[pass]
		var totalTemp int = hitungTotalHargaPesanan(temp.Items, temp.JumlahItems)
		for i > 0 {
			var totalPrev int = hitungTotalHargaPesanan(dataUrut[i-1].Items, dataUrut[i-1].JumlahItems)
			if totalTemp < totalPrev {
				dataUrut[i] = dataUrut[i-1]
				i--
			} else {
				break
			}
		}
		dataUrut[i] = temp
		pass++
	}

	fmt.Print("Masukkan total harga pesanan yang dicari: ")
	var totalCari int = bacaInt()

	var left int = 0
	var right int = jumlahPesananSaatIni - 1
	var found bool = false
	var foundIndex int = -1

	for left <= right {
		var mid int = (left + right) / 2 
		var totalMid int = hitungTotalHargaPesanan(dataUrut[mid].Items, dataUrut[mid].JumlahItems)

		if totalMid == totalCari {
			found = true
			foundIndex = mid
			break
		} else if totalMid < totalCari {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if found {
		messageTeks[messageCount] = fmt.Sprintf("Pesanan dengan total harga Rp%d ditemukan. ID Pesanan: %d.", totalCari, dataUrut[foundIndex].ID)
		messageCount++
	} else {
		messageTeks[messageCount] = fmt.Sprintf("Pesanan dengan total harga Rp%d tidak ditemukan.", totalCari)
		messageCount++
	}
	cetakDenganBingkai(messageTeks, messageCount)
}