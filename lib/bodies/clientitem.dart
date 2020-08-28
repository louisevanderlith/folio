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

    _loaded = txtName != null
      && txtSecret != null
      && txtUrl != null
      && lstResources != null
      && chkCodes != null
      && chkTerms != null;
  }

  bool get loaded {
    return _loaded;
  }

  String get name {
    return txtName.text;
  }

  String get secret {
    return txtSecret.text;
  }

  String get url {
    return txtUrl.text;
  }

  List<String> get resources {
    return null; //lstResources
  }

  bool get terms {
    return chkTerms.checked;
  }

  bool get codes {
    return chkCodes.checked;
  }

  Client toDTO() {
    return new Client(name, secret, url, resources, terms, codes);
  }
}
