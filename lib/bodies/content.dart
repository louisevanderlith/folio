import 'banner.dart';
import 'colour.dart';
import 'information.dart';
import 'section.dart';

class Content {
  final String realm;
  final String client;
  final String language;
  final String email;
  final Banner banner;
  final Section sectionA;
  final Section sectionB;
  final Information info;
  final Colour colour;

  Content(this.realm, this.client, this.language, this.email, this.banner,
      this.sectionA, this.sectionB, this.info, this.colour);

  Map<String, dynamic> toJson() {
    return {
      "Realm": realm,
      "Client": client,
      "Language": language,
      "Email": email,
      "Banner": banner,
      "SectionA": sectionA,
      "SectionB": sectionB,
      "Info": info,
      "Colour": colour,
    };
  }
}
