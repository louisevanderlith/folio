import 'dart:convert';
import 'dart:html';

import 'package:mango_secure/bodies/profile.dart';
import 'package:mango_secure/profileapi.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';

import 'basicsite.dart';
import 'clientsform.dart';
import 'contactsform.dart';

class ProfileForm extends FormState {
  Key _objKey;

  BasicSiteForm basicForm;
  ContactsForm contactsForm;
  ClientsForm clientsForm;

  ProfileForm(Key k) : super("#frmProfile", "#btnSubmit") {
    _objKey = k;
    basicForm = new BasicSiteForm();
    contactsForm = new ContactsForm();
    clientsForm = new ClientsForm();
  }

  void onSubmitClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);
      final obj = new Profile(
          basicForm.title,
          basicForm.description,
          contactsForm.items,
          basicForm.imageKey,
          clientsForm.items,
          basicForm.endpoints,
          basicForm.codes,
          basicForm.terms);

      HttpRequest req;
      if (_objKey.toJson() != "0`0") {
        req = await updateProfile(_objKey, obj);
      } else {
        req = await createProfile(obj);
      }

      var result = jsonDecode(req.response);

      if (req.status == 200) {
        final data = result['Data'];
        final rec = data['Record'];

        if (rec != null) {
          final key = rec['K'];

          _objKey = key;
        }
      }
    }
  }
}
