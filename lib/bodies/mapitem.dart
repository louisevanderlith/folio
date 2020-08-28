import 'dart:html';

class MapItem {
  TextInputElement txtName;
  TextInputElement txtValue;
  bool _loaded;

  MapItem(String nameId, String valueId) {
    txtName = querySelector(nameId);
    txtValue = querySelector(valueId);

    _loaded = txtName != null && txtValue != null;
  }

  bool get loaded {
    return _loaded;
  }

  String get name {
    return txtName.text;
  }

  String get value {
    return txtValue.text;
  }

  MapEntry toEntry() {
    return new MapEntry(this.name, this.value);
  }
}
