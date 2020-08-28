import 'package:FOLIO.APP/profileform.dart';
import 'package:mango_ui/keys.dart';

void main() async {
  final k = getObjKey();
  print("KEY: ${k.toString()}");
  new ProfileForm(k);
}
