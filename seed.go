package main

import (
    "log"

    "github.com/joho/godotenv"
    "github.com/annddvaa/gin-firebase-backend/config"
    "github.com/annddvaa/gin-firebase-backend/models"
    "gorm.io/gorm"
)

func main() {
    godotenv.Load()
    config.InitDatabase()

    // Hapus data lama agar bersih
    config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Product{})

    products := []models.Product{
        // --- Layar ---
        {Name: "Ganti LCD iPhone 13 Pro Max", Price: 3500000, Category: "Layar", Stock: 10, Description: "Layar original OLED dengan true tone aktif", ImageURL: "https://picsum.photos/seed/layar1/400/400"},
        {Name: "Ganti LCD iPhone 11", Price: 850000, Category: "Layar", Stock: 15, Description: "Layar original copotan bergaransi", ImageURL: "https://picsum.photos/seed/layar2/400/400"},
        {Name: "Ganti LCD Samsung S23 Ultra", Price: 4200000, Category: "Layar", Stock: 5, Description: "Layar dynamic AMOLED 2X original", ImageURL: "https://picsum.photos/seed/layar3/400/400"},
        {Name: "Ganti LCD Samsung A54", Price: 1200000, Category: "Layar", Stock: 12, Description: "Layar super AMOLED original pabrik", ImageURL: "https://picsum.photos/seed/layar4/400/400"},
        {Name: "Ganti LCD Xiaomi Redmi Note 12", Price: 650000, Category: "Layar", Stock: 20, Description: "Layar original OLED bergaransi", ImageURL: "https://picsum.photos/seed/layar5/400/400"},
        {Name: "Ganti LCD Poco X5", Price: 750000, Category: "Layar", Stock: 10, Description: "Layar Amoled 120hz original", ImageURL: "https://picsum.photos/seed/layar6/400/400"},
        {Name: "Ganti LCD Oppo Reno 8", Price: 1100000, Category: "Layar", Stock: 8, Description: "Layar original amoled termasuk pasang", ImageURL: "https://picsum.photos/seed/layar7/400/400"},
        {Name: "Ganti LCD Vivo V27", Price: 1350000, Category: "Layar", Stock: 5, Description: "Layar lengkung amoled original", ImageURL: "https://picsum.photos/seed/layar8/400/400"},
        {Name: "Ganti Kaca Depan iPhone 12", Price: 500000, Category: "Layar", Stock: 20, Description: "Hanya ganti kaca depan (LCD harus normal)", ImageURL: "https://picsum.photos/seed/layar9/400/400"},
        {Name: "Ganti Kaca Depan Samsung S22", Price: 650000, Category: "Layar", Stock: 15, Description: "Ganti kaca depan original corning", ImageURL: "https://picsum.photos/seed/layar10/400/400"},

        // --- Baterai ---
        {Name: "Ganti Baterai iPhone 13 Pro", Price: 650000, Category: "Baterai", Stock: 20, Description: "Baterai original health 100%", ImageURL: "https://picsum.photos/seed/bat1/400/400"},
        {Name: "Ganti Baterai iPhone X", Price: 350000, Category: "Baterai", Stock: 30, Description: "Baterai original health 100% no minus", ImageURL: "https://picsum.photos/seed/bat2/400/400"},
        {Name: "Ganti Baterai Samsung S21", Price: 450000, Category: "Baterai", Stock: 15, Description: "Baterai original SEIN garansi 1 bulan", ImageURL: "https://picsum.photos/seed/bat3/400/400"},
        {Name: "Ganti Baterai Poco F3", Price: 250000, Category: "Baterai", Stock: 25, Description: "Baterai original fast charging aman", ImageURL: "https://picsum.photos/seed/bat4/400/400"},
        {Name: "Ganti Baterai Xiaomi Redmi Note 10", Price: 200000, Category: "Baterai", Stock: 40, Description: "Baterai double power garansi 3 bulan", ImageURL: "https://picsum.photos/seed/bat5/400/400"},
        {Name: "Ganti Baterai Oppo A5s", Price: 180000, Category: "Baterai", Stock: 50, Description: "Baterai tahan lama kualitas A", ImageURL: "https://picsum.photos/seed/bat6/400/400"},
        {Name: "Ganti Baterai Vivo Y12", Price: 180000, Category: "Baterai", Stock: 35, Description: "Baterai baru kualitas prima", ImageURL: "https://picsum.photos/seed/bat7/400/400"},
        {Name: "Ganti Baterai iPhone 14", Price: 850000, Category: "Baterai", Stock: 10, Description: "Baterai original bawaan apple 100%", ImageURL: "https://picsum.photos/seed/bat8/400/400"},
        {Name: "Ganti Baterai iPad Air 4", Price: 950000, Category: "Baterai", Stock: 5, Description: "Baterai tablet besar tahan lama", ImageURL: "https://picsum.photos/seed/bat9/400/400"},
        {Name: "Ganti Baterai Samsung Tab S7", Price: 850000, Category: "Baterai", Stock: 8, Description: "Baterai tablet original", ImageURL: "https://picsum.photos/seed/bat10/400/400"},

        // --- Software ---
        {Name: "Flashing/Install Ulang Android", Price: 100000, Category: "Software", Stock: 999, Description: "Flashing rom global resmi", ImageURL: "https://picsum.photos/seed/soft1/400/400"},
        {Name: "Restore iOS & Update", Price: 120000, Category: "Software", Stock: 999, Description: "Clean install iOS terbaru bebas bug", ImageURL: "https://picsum.photos/seed/soft2/400/400"},
        {Name: "Unlock Pola/Sandi Android", Price: 150000, Category: "Software", Stock: 999, Description: "Buka kunci layar tanpa hapus data (jika didukung)", ImageURL: "https://picsum.photos/seed/soft3/400/400"},
        {Name: "Bypass FRP Google Account", Price: 200000, Category: "Software", Stock: 999, Description: "Bypass akun google terkunci", ImageURL: "https://picsum.photos/seed/soft4/400/400"},
        {Name: "Unlock Mi Cloud", Price: 250000, Category: "Software", Stock: 999, Description: "Hapus akun mi cloud permanen", ImageURL: "https://picsum.photos/seed/soft5/400/400"},
        {Name: "Fix Bootloop Xiaomi", Price: 150000, Category: "Software", Stock: 999, Description: "Perbaikan hp mentok di logo mi", ImageURL: "https://picsum.photos/seed/soft6/400/400"},
        {Name: "Fix Bootloop Samsung", Price: 150000, Category: "Software", Stock: 999, Description: "Perbaikan hp mentok di logo samsung", ImageURL: "https://picsum.photos/seed/soft7/400/400"},
        {Name: "Instalasi Aplikasi Premium", Price: 50000, Category: "Software", Stock: 999, Description: "Paket aplikasi premium (spotify, netflix, dll)", ImageURL: "https://picsum.photos/seed/soft8/400/400"},
        {Name: "Backup & Restore Data", Price: 100000, Category: "Software", Stock: 999, Description: "Pindah data aman ke hp baru 100%", ImageURL: "https://picsum.photos/seed/soft9/400/400"},
        {Name: "Fix Stuck Logo Apple (iTunes)", Price: 150000, Category: "Software", Stock: 999, Description: "Perbaikan error itunes stuck logo", ImageURL: "https://picsum.photos/seed/soft10/400/400"},

        // --- Kamera ---
        {Name: "Ganti Kamera Utama iPhone 13", Price: 1250000, Category: "Kamera", Stock: 5, Description: "Kamera utama bening auto fokus jalan", ImageURL: "https://picsum.photos/seed/cam1/400/400"},
        {Name: "Ganti Kamera Depan iPhone 11", Price: 650000, Category: "Kamera", Stock: 8, Description: "Kamera depan jernih", ImageURL: "https://picsum.photos/seed/cam2/400/400"},
        {Name: "Ganti Kamera Utama Samsung S22", Price: 950000, Category: "Kamera", Stock: 5, Description: "Kamera utama jernih", ImageURL: "https://picsum.photos/seed/cam3/400/400"},
        {Name: "Ganti Kaca Kamera iPhone 14 Pro", Price: 350000, Category: "Kamera", Stock: 15, Description: "Kaca sapphire original tidak mudah baret", ImageURL: "https://picsum.photos/seed/cam4/400/400"},
        {Name: "Fix Kamera Blur Poco X3", Price: 450000, Category: "Kamera", Stock: 10, Description: "Perbaikan lensa kamera gagal fokus", ImageURL: "https://picsum.photos/seed/cam5/400/400"},
        {Name: "Ganti Kamera Depan Oppo F11", Price: 350000, Category: "Kamera", Stock: 12, Description: "Modul kamera pop-up original", ImageURL: "https://picsum.photos/seed/cam6/400/400"},
        {Name: "Ganti Modul Kamera Vivo V20", Price: 400000, Category: "Kamera", Stock: 10, Description: "Kamera auto fokus normal kembali", ImageURL: "https://picsum.photos/seed/cam7/400/400"},
        {Name: "Fix Face ID iPhone X - 12", Price: 650000, Category: "Kamera", Stock: 999, Description: "Reparasi modul dot projector face id", ImageURL: "https://picsum.photos/seed/cam8/400/400"},
        {Name: "Fix Face ID iPhone 13 - 14", Price: 850000, Category: "Kamera", Stock: 999, Description: "Reparasi face id true depth camera", ImageURL: "https://picsum.photos/seed/cam9/400/400"},
        {Name: "Ganti Kaca Kamera Samsung S23", Price: 250000, Category: "Kamera", Stock: 20, Description: "Ganti kaca luar kamera retak", ImageURL: "https://picsum.photos/seed/cam10/400/400"},

        // --- Charging ---
        {Name: "Ganti Port Charger iPhone (Lightning)", Price: 350000, Category: "Charging", Stock: 20, Description: "Ganti flex charger original koneksi lancar", ImageURL: "https://picsum.photos/seed/char1/400/400"},
        {Name: "Ganti Port Charger Type-C Samsung", Price: 250000, Category: "Charging", Stock: 30, Description: "Support fast charging original", ImageURL: "https://picsum.photos/seed/char2/400/400"},
        {Name: "Ganti Port Charger Type-C Xiaomi", Price: 150000, Category: "Charging", Stock: 40, Description: "Board charger original turbo charge jalan", ImageURL: "https://picsum.photos/seed/char3/400/400"},
        {Name: "Fix IC Charger iPhone 11", Price: 750000, Category: "Charging", Stock: 999, Description: "Perbaikan mesin ic usb/charger", ImageURL: "https://picsum.photos/seed/char4/400/400"},
        {Name: "Fix IC Charger Android", Price: 450000, Category: "Charging", Stock: 999, Description: "Perbaikan mesin hp tidak bisa di cas", ImageURL: "https://picsum.photos/seed/char5/400/400"},
        {Name: "Ganti Port Charger Micro USB", Price: 100000, Category: "Charging", Stock: 50, Description: "Ganti konektor cas lama", ImageURL: "https://picsum.photos/seed/char6/400/400"},
        {Name: "Fix Wireless Charging iPhone", Price: 550000, Category: "Charging", Stock: 10, Description: "Ganti koil wireless charging magsafe", ImageURL: "https://picsum.photos/seed/char7/400/400"},
        {Name: "Servis Jalur Cas Putus", Price: 350000, Category: "Charging", Stock: 999, Description: "Jumper jalur konektor cas di mesin", ImageURL: "https://picsum.photos/seed/char8/400/400"},
        {Name: "Ganti Flex Charger iPad", Price: 650000, Category: "Charging", Stock: 5, Description: "Ganti port cas khusus ipad type c/lightning", ImageURL: "https://picsum.photos/seed/char9/400/400"},
        {Name: "Fix Fast Charging Tidak Fungsi", Price: 250000, Category: "Charging", Stock: 999, Description: "Perbaikan error lambat ngecas", ImageURL: "https://picsum.photos/seed/char10/400/400"},

        // --- Lainnya ---
        {Name: "Servis Mesin Mati Total", Price: 1250000, Category: "Lainnya", Stock: 999, Description: "Perbaikan hp mati mendadak (harga mulai dari)", ImageURL: "https://picsum.photos/seed/oth1/400/400"},
        {Name: "Servis IC Audio iPhone", Price: 850000, Category: "Lainnya", Stock: 999, Description: "Fix speaker mati, mic mati (audio codec)", ImageURL: "https://picsum.photos/seed/oth2/400/400"},
        {Name: "Servis Sinyal Hilang (No Service)", Price: 750000, Category: "Lainnya", Stock: 999, Description: "Perbaikan ic baseband/WTR sinyal", ImageURL: "https://picsum.photos/seed/oth3/400/400"},
        {Name: "Ganti Backglass iPhone 12/13/14", Price: 650000, Category: "Lainnya", Stock: 20, Description: "Ganti kaca belakang rapi dengan laser", ImageURL: "https://picsum.photos/seed/oth4/400/400"},
        {Name: "Ganti Tombol Power/Volume", Price: 250000, Category: "Lainnya", Stock: 30, Description: "Ganti flexible tombol keras/tidak fungsi", ImageURL: "https://picsum.photos/seed/oth5/400/400"},
        {Name: "Pembersihan Water Damage", Price: 350000, Category: "Lainnya", Stock: 999, Description: "Pembersihan karat dan short akibat air", ImageURL: "https://picsum.photos/seed/oth6/400/400"},
        {Name: "Ganti Speaker Bawah (Buzzer)", Price: 250000, Category: "Lainnya", Stock: 20, Description: "Speaker musik pecah/sember/mati", ImageURL: "https://picsum.photos/seed/oth7/400/400"},
        {Name: "Ganti Earpiece (Speaker Telepon)", Price: 200000, Category: "Lainnya", Stock: 20, Description: "Suara telepon kecil/tidak terdengar", ImageURL: "https://picsum.photos/seed/oth8/400/400"},
        {Name: "Servis Mic Mati / Suara Kecil", Price: 200000, Category: "Lainnya", Stock: 25, Description: "Lawan bicara tidak bisa mendengar suara", ImageURL: "https://picsum.photos/seed/oth9/400/400"},
        {Name: "Ganti Housing / Frame Fullset", Price: 850000, Category: "Lainnya", Stock: 10, Description: "Ganti tulang hp bengkok/penyok jadi baru", ImageURL: "https://picsum.photos/seed/oth10/400/400"},
    }

    for _, p := range products {
        config.DB.Create(&p)
    }
    log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}
