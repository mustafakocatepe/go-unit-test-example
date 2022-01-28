Go standard Testing Package 
https://go.dev/doc/faq#testing_framework


Go’da test kodlarımızın bulunduğu dosyalar, testi yazılan dosyanın sonuna “_test” eklenerek isimlendirilir. Örneğin control_test.go
Go’da test fonksiyonunun ismi, testi yapılan fonksiyon isminin başına “Test” eklenerek oluşturulur. Örneğin TestSum.
Test fonksiyonları (TestTopla gibi) her zaman sadece t *testing.T parametresini alır ve herhangi bir şey döndürmez. Test sonucu doğrudan t.Error() ile ekrana yazdırılır.

Test parametreleri için yardım komutu
➜ go help testflag

Belirli bir klasördeki testleri çalıştırma
➜ go test ./control/control_test -v

Test Coverage
➜ go test -coverprofile cover.out
