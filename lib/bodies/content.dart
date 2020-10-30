import 'package:mango_ui/keys.dart';

import 'banner.dart';
import 'colour.dart';
import 'contact.dart';
import 'information.dart';
import 'section.dart';

class Content {
  final String realm;
  final String client;
  final Key logo;
  final String language;
  final String email;
  final Banner banner;
  final Section sectionA;
  final Section sectionB;
  final Information info;
  final Colour colour;
  final List<Contact> contacts;

  Content(
      this.realm,
      this.client,
      this.logo,
      this.language,
      this.email,
      this.banner,
      this.sectionA,
      this.sectionB,
      this.info,
      this.colour,
      this.contacts);

  Map<String, dynamic> toJson() {
    return {
      "Realm": realm,
      "Client": client,
      "LogoKey": logo,
      "Language": language,
      "Banner": banner,
      "SectionA": sectionA,
      "SectionB": sectionB,
      "Info": info,
      "Colour": colour,
      "Email": email,
      "Contacts": contacts
    };
  }
}
