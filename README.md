# <p align="center">PulseFin || Financial Dashboard </p>

<br>


## Proje Amacı
* Bu proje, kullanıcıların kripto para piyasalarındaki güncel fiyat değişimlerini, işlem hacimlerini ve piyasa trendlerini görselleştirmelerine olanak tanıyan bir Financial Dashboard uygulamasının frontend kısmıdır. Ayrıca, piyasaya dair önemli haberleri kullanıcı dostu bir arayüzle sunarak bilgilendirme sağlar.

<br>

## Proje Hedefleri

1) Kişiselleştirilmiş Takip: Kullanıcıların izleme listesine kendi belirledikleri kripto paraları ekleyebilmesi.
2) Gerçek Zamanlı Veri: İzleme listesindeki kripto paraların fiyatlarını, günlük değişimlerini ve işlem hacimlerini WebSocket ile gerçek zamanlı olarak sunmak.
3) Bilgilendirme: Kripto paralara dair önemli haberleri modern ve erişilebilir bir arayüzle kullanıcılara sunmak.
4) Gelecekteki Genişletme: Kullanıcıların cüzdanlarını bağlayarak gelir-gider dengesini ve yatırım performanslarını analiz edebilecekleri bir yapı oluşturmak.


<br>

## Projenin Özellikleri

**1. Heatmap Görselleştirmesi**

- İzleme listesine eklenen kripto paraların fiyat ve hacim verileri renk kodlu bir heatmap üzerinde görselleştirilir.
- Heatmap, kullanıcıya yatırım portföyündeki varlıkların performansını hızlı ve görsel bir şekilde analiz etme imkânı sunar.

- Her hücrede şu bilgiler gösterilir:
`Kripto para adı (ör. BTC-USDT)`
`Son fiyat`
`Günlük hacim`
`Günlük değişim yüzdesi`


<br>


**2. Haber Listesi**
- Kripto paralarla ilgili güncel ve önemli haberleri kullanıcıya sunar.
- Haber başlıklarına tıklanarak detaylı bilgiye ulaşılabilir.
- Kullanıcılar, izledikleri varlıklarla ilgili gelişmeleri takip ederek daha bilinçli kararlar alabilir.

<br>


**3. İzleme Listesi Yönetimi**
- Kullanıcılar, kendi belirledikleri kripto paraları izleme listesine ekleyebilir veya çıkarabilir.
- İzleme listesindeki varlıkların fiyatları ve değişimleri gerçek zamanlı olarak güncellenir.

<br>


**4. Cüzdan Entegrasyonu (Gelecekte Eklenecek)**
- Kullanıcılar, kripto para cüzdanlarını uygulamaya bağlayarak gelir-gider dengesini ve yatırım performanslarını görebilecekler.
- Bu özellik, kullanıcının tüm finansal durumunu tek bir ekranda yönetmesine olanak tanıyacak.

<br>



## Kulanım Kılavuzu

### **Projeyi Çalıştırma**
1. **Bağımlılıkları Yükleme :**
   ```bash
    go mod tidy
   ```

2. **Çevresel Değişkenleri Ayarlama :**
   ```bash
    SERVER_PORT=8080
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=financial_dashboard
    DB_SSLMODE=disable
    
    JWT_SECRET=your_jwt_secret
    
    EMAIL_HOST=smtp.your-email-provider.com
    EMAIL_PORT=587
    EMAIL_USER=your_email@example.com
    EMAIL_PASSWORD=your_email_password
    EMAIL_FROM=Financial Dashboard <your_email@example.com>

   ```
3. **Uygulamayı Başlatma**
   ```bash
    go run main.go
   ```
   
4. **Frontend**
* Bu backend uygulaması, kripto para fiyatlarını ve haberleri sağlamak için bir frontend uygulaması ile entegre çalışır. Frontend, kullanıcıların izleme listelerini yönetmelerini ve fiyat/haber verilerini görselleştirmelerini sağlar.

* [Frontend GitHub Deposu](https://github.com/Efe-Eroglu/financial-dashboard.git)



<br>


## Katkıda Bulunma
* Projede bir hata bulursanız veya bir geliştirme için pull request açabilirsizin.
