#### Problem 3
##### Compare Folder

TASK #1. Implementasikan sebuah *program* yang membandingkan isi dari dua direktori melalui *parameter*.
1. Jika file ada di source dan target, abaikan
2. Jika file ada di source tapi tidak ada di target berikan keterangan **NEW**
3. Jika file tidak ada di source tapi ada di target berikan keterangan **DELETED**

ILUSTRASI: membandingkan *source* direktori terhadap *target* direktori
> $ compare source/ target/

```
source
│   README.md               
│   file001.txt    
│
└───folder1
│   │   file011.txt
│   │   file012.txt
│   │
│   └───subfolder1
│       │   file111.txt
│       │   file112.txt       
│   
└───folder2
    │   file021.txt
    │   file022.txt
```

```
target
│   README.md
│   file001.txt    
│
└───folder1
│   │   file011.txt
│   │   file012.txt
│   │   file013.txt
│   │
│   └───subfolder1
│       │   file111.txt
│       │   file113.txt       
│   
└───folder2
    │   file021.txt
```

OUPUT:
```
folder1/file013.txt DELETED
folder1/subfolder1/file112.txt NEW
folder1/subfolder1/file113.txt DELETED
folder2/file022.txt NEW

```

TASK #2. Modifikasi program #1 untuk *compare* file content untuk rule (1), jika ada perbedaan beri keterangan **MODIFIED** 
