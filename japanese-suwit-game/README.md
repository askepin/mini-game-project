# japanese-suwit-game
just for fun project when i learn Golang to motivate me for better every single day

in Bahasa Indonesia

Kode ini adalah sebuah program sederhana dalam bahasa Go yang mengimplementasikan permainan batu gunting kertas antara pemain dan komputer.

Berikut ini adalah penjelasan singkat mengenai kode tersebut:

1. Impor Paket:
		- 'fmt' digunakan untuk input/output standar.
		- 'math/rand' digunakan untuk menghasilkan angka acak.
		- 'strings' digunakan untuk operasi terkait string.
		- 'time' digunakan untuk menghasilkan nilai seed yang unik pada 'rand.Seed()'.

2. Fungsi 'determineWinner':

		- Menerima dua argumen 'playerChoice' dan 'computerChoice' yang merupakan pilihan pemain dan pilihan komputer.
		- Menggunakan serangkaian kondisi untuk menentukan pemenang.
		- Mengembalikan string "Seri" jika kedua pilihan sama, "Anda Menang!" jika pemain menang, dan "Computer Menang!" jika komputer menang.

3. Fungsi 'main':

		- Mengatur inisialisasi 'seed' untuk pengacakan angka acak menggunakan 'rand.Seed(time.Now().UnixNano())'.
		- Membuat slice 'choices' yang berisi pilihan yang mungkin, yaitu "batu", "gunting", dan "kertas".
		- Menggunakan variabel 'playAgain' untuk mengontrol apakah permainan akan diulang atau tidak.
		- Menginisialisasi variabel 'playerScore' dan 'computerScore' dengan nilai awal 0 untuk menyimpan skor pemain dan komputer
		
		Di dalam perulangan 'for', langkah-langkah berikut terjadi:

		- Mencetak instruksi untuk pemain memilih "batu", "gunting", atau "kertas".
		- Membaca input pemain menggunakan 'fmt.Scanln()' dan menyimpannya ke variabel 'playerChoice'.
		- Mengubah 'playerChoice' menjadi huruf kecil menggunakan 'strings.ToLower()'.
		- Memvalidasi pilihan pemain dengan memeriksa apakah 'playerChoice' terdapat dalam 'choices'.
		- Jika pilihan pemain tidak valid, mencetak pesan kesalahan dan melanjutkan perulangan.
		- Menghasilkan pilihan komputer secara acak dari 'choices' menggunakan 'rand.Intn(len(choices))'.
		- Mengubah 'computerChoice' menjadi huruf kecil menggunakan 'strings.ToLower()'.
		- Mencetak pilihan pemain dan komputer.
		- Memanggil fungsi 'determineWinner()' untuk menentukan pemenang.
		- Mencetak hasilnya.
		- Jika pemain menang, skor pemain 'playerScore' akan ditambah 1.
		- Jika komputer menang, skor komputer 'computerScore' akan ditambah 1.
		- Jika tidak ada pemenang, cetak pesan bahwa tidak ada penambahan skor.
		- Mencetak skor pemain dan skor komputer
		- Membaca input dari pemain untuk menentukan apakah permainan akan diulang atau tidak.
		- Mengubah 'playAgainInput' menjadi huruf kecil menggunakan 'strings.ToLower()'.
		- Jika 'playAgainInput' bukan "ya", mengubah nilai 'playAgain' menjadi 'false', sehingga keluar dari perulangan dan program selesai.
		
Demikianlah penjelasan mengenai kode yang Anda berikan. Program ini memungkinkan pemain untuk bermain permainan batu gunting kertas melawan komputer dan melacak skor pemain dan skor komputer selama permainan. Setelah setiap putaran, skor akan diperbarui dan ditampilkan kepada pemain.
