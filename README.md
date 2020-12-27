# Metode SAW (Simple Additive Weighting)
Metode Simple Additive Weighting (SAW) dikenal dengan istilah metode penjumlahan terbobot. Konsep dasar pada metode SAW adalah mencari penjumlahan terbobot dari rating kinerja pada setiap alternatif di semua atribut

## Studi kasus
Menentukan rekomendasi tempat bootcamp IT
kriteria yang digunakan :

1. C1 = Biaya (seberapa banyak biaya untuk bootcamp)
2. C2 = Lokasi (seberapa jauh lokasi tempat bootcamp dari rumah)
3. C3 = Fasilitas ( fasilitas tempat bootcamp)
4. C4 = Kualitas pengajar (dilihat dari pengalaman para mentor di tempat bootcamp)
dari kriteria tersebut akan digunakan variable. dimana setiap variabel akan diberi bobot dalam bentuk angka dengan 
range **1-5** 
---
| No.  | Biaya  | Nilai |
| :------------ |:---------------:| -----:|
| 1. | dibawah 10.000.000 | 2 |
| 2. | 10.000.000 - 20.000.000 | 3 |
| 3. | 20.000.000 - 25.000.000 | 4 |
| 4. | diatas 25.000.000 | 5 |
---
