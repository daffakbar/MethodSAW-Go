# Metode SAW (Simple Additive Weighting)
Metode Simple Additive Weighting (SAW) dikenal dengan istilah metode penjumlahan terbobot. Konsep dasar pada metode SAW adalah mencari penjumlahan terbobot dari rating kinerja pada setiap alternatif di semua atribut

## Studi kasus
Menentukan rekomendasi tempat bootcamp IT
kriteria yang digunakan :

1. C1 = Biaya (seberapa banyak biaya untuk bootcamp)
2. C2 = Lokasi (seberapa jauh lokasi tempat bootcamp dari rumah)
3. C3 = Fasilitas ( fasilitas tempat bootcamp)
4. C4 = Kualitas pengajar (dilihat dari pengalaman para mentor di tempat bootcamp)

Bobot Kriteria
| Kriteria  | Bobot  |
| :------------ |---------------:| 
| C1. | 0,4 | 
| C2. | 0,3 | 
| C3. | 0,2 | 
| C4. | 0,1 | 
| TOTAL. | 1 | 
adasdasd
asdasd
sda
dari kriteria tersebut akan digunakan variable. dimana setiap variabel akan diberi bobot dalam bentuk angka dengasn range 1 - 5 
---
#### Biaya
| No.  | Biaya  | Nilai |
| :------------ |:---------------:| -----:|
| 1. | dibawah 10.000.000 | 2 |
| 2. | 10.000.000 - 20.000.000 | 3 |
| 3. | 20.000.000 - 25.000.000 | 4 |
| 4. | diatas 25.000.000 | 5 |

#### Lokasi
| No.  | Lokasi  | Nilai |
| :------------ |:---------------:| -----:|
| 1. | dibawah 5 km | 2 |
| 2. | 5 - 10 km | 3 |
| 3. | 10 -15 km | 4 |
| 4. | diatas 15km | 5 |

#### Fasilitas
| No.  | Fasilitas  | Nilai |
| :------------ |:---------------:| -----:|
| 1. | Tidak lengkap | 1 |
| 2. | Cukup lengkap | 3 |
| 3. | Sangat lengkap | 5 |

#### "Kualitas pengajar
| No.  | Kualitas pengajar  | Nilai |
| :------------ |:---------------:| -----:|
| 1. | pengalaman kerja 2 tahun | 3 |
| 2. | pengalaman kerja 2 - 3 tahun | 4 |
| 3. | pengalaman kerja 3 tahun lebih | 5 |

---
### Perhitungan
pada studi kasus ini, ada 3 tempat bimbingan belajar yang akan menjadi alternatif, yaitu:
1. Arkademy (A1)
2. Hactiv8 (A2)
3. Dumbways (A3)
setiap alternatif diberikan variable untuk masing - masing kriteria sesuai keadaan dari alternatif: 
| Alternatif  | C1  | C2 | C3 | C4
| :------------ |:---------------:|:---------------:|:---------------:| -----:|
| A1. | 3 | 2 | 3 | 4 |
| A2. | 2 | 4 | 3 | 3 |
| A3. | 4 | 5 | 5 | 5 |
