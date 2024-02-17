<!-- customize-category:Java -->

# JDK bin tools

## keytool

将证书和私钥导出为 PKCS#12 文件，PKCS#12 文件将同时包含证书和私钥

```sh
openssl pkcs12 -export -clcerts -in lance-zheng.cer -inkey rsa_private.key -out client.p12 -passout pass:123456
```

将 PKCS#12 文件导入到 JKS 文件中，

```sh
keytool -importkeystore -srckeystore client.p12 -srcstoretype pkcs12 -destkeystore clientkeystore.jks -deststoretype JKS
```

通过上面的证书进行

```java
import java.io.FileInputStream;
import java.net.URL;
import java.security.KeyStore;
import javax.net.ssl.*;

public class HttpsClientExample {
    public static void main(String[] args) throws Exception {
        // Load the Java keystore (JKS) containing the client certificate and private key
        char[] keystorePassword = "your_keystore_password".toCharArray();
        String keystoreFile = "path/to/clientkeystore.jks";
        KeyStore keyStore = KeyStore.getInstance("JKS");
        FileInputStream fis = new FileInputStream(keystoreFile);
        keyStore.load(fis, keystorePassword);
        fis.close();

        // Create a KeyManagerFactory with the loaded keystore
        KeyManagerFactory keyManagerFactory = KeyManagerFactory.getInstance(KeyManagerFactory.getDefaultAlgorithm());
        keyManagerFactory.init(keyStore, keystorePassword);

        // Create an SSLContext with the KeyManagers from the keystore
        SSLContext sslContext = SSLContext.getInstance("TLS");
        sslContext.init(keyManagerFactory.getKeyManagers(), null, null);

        // Set the SSLContext as default for HTTPS connections
        HttpsURLConnection.setDefaultSSLSocketFactory(sslContext.getSocketFactory());

        // Now you can make HTTPS requests using HttpsURLConnection or HttpClient
        // For example:
        URL url = new URL("https://example.com");
        HttpsURLConnection connection = (HttpsURLConnection) url.openConnection();
        // ... (Set request properties, handle response, etc.)
    }
}

```
