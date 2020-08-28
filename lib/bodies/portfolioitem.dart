// import 'dart:html';

// import 'package:mango_ui/bodies/portfolio.dart';

// class PortfolioItem {
//   TextInputElement _icon;
//   TextInputElement _name;
//   TextInputElement _descr;
//   UrlInputElement _url;
//   bool _loaded;

//   PortfolioItem(
//       String iconElem, String nameElem, String descrElem, String urlElem) {
//     _icon = querySelector(iconElem);
//     _name = querySelector(nameElem);
//     _descr = querySelector(descrElem);
//     _url = querySelector(urlElem);

//     _loaded = _icon != null && _name != null && _descr != null && _url != null;
//   }

//   String get icon {
//     return _icon.value;
//   }

//   String get name {
//     return _name.value;
//   }

//   String get url {
//     return _url.value;
//   }

//   String get description {
//     return _descr.value;
//   }

//   bool loaded() {
//     return _loaded;
//   }

//   Portfolio toDTO() {
//     return new Portfolio(icon, url, name, description);
//   }
// }
