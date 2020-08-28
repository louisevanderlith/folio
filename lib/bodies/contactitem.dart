import 'dart:html';

import 'package:mango_secure/bodies/contact.dart';

class ContactItem {
  TextInputElement txtName;
  TextInputElement txtIcon;
  TextInputElement txtValue;

  bool _loaded;

  ContactItem(String nameId, String iconId, String valueId) {
    txtName = querySelector(nameId);
    txtIcon = querySelector(iconId);
    txtValue = querySelector(valueId);

    _loaded = txtName != null && txtIcon.value != null && txtValue != null;
  }

  bool get loaded {
    return _loaded;
  }

  String get name {
    return txtName.text;
  }

  String get icon {
    return txtIcon.text;
  }

  String get value {
    return txtValue.text;
  }

  Contact toDTO() {
    return new Contact(icon, name, value);
  }
}
