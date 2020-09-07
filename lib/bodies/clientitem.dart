import 'dart:html';

import 'package:mango_secure/bodies/client.dart';

class ClientItem {
  TextInputElement txtName;
  PasswordInputElement txtSecret;
  UrlInputElement txtUrl;
  UListElement lstResources;
  CheckboxInputElement chkCodes;
  CheckboxInputElement chkTerms;

  bool _loaded;

  ClientItem(String nameId, String secretId, String urlId, String resourcesId,
      String codesId, String termsId) {
    txtName = querySelector(nameId);
    txtSecret = querySelector(secretId);
    txtUrl = querySelector(urlId);
    lstResources = querySelector(resourcesId);
    chkCodes = querySelector(codesId);
    chkTerms = querySelector(termsId);

    _loaded = txtName != null &&
        txtSecret != null &&
        txtUrl != null &&
        lstResources != null &&
        chkCodes != null &&
        chkTerms != null;
  }

  bool get loaded {
    return _loaded;
  }

  String get name {
    return txtName.value;
  }

  String get secret {
    return txtSecret.value;
  }

  String get url {
    return txtUrl.value;
  }

  List<String> get resources {
    return lstResources.children.map((e) {
      print(e);
      return e.text;
    });
  }

  bool get terms {
    return chkTerms.checked;
  }

  bool get codes {
    return chkCodes.checked;
  }

  Client toDTO() {
    print("Client Name ${name}");
    print("Client Url ${url}");
    print("Client Resources ${resources}");
    print("Client Terms ${terms}");
    print("Client Codes ${codes}");
    return new Client(name, secret, url, resources, terms, codes);
  }
}
