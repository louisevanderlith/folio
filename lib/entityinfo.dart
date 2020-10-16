import 'dart:html';

import 'package:mango_entity/bodies/address.dart';
import 'package:mango_entity/bodies/entity.dart';
import 'package:mango_secure/bodies/mapitem.dart';
import 'package:mango_ui/keys.dart';

import 'addressform.dart';

class EntityInfoForm {
  TextInputElement txtName;
  SelectElement ddlProfileKey;
  TextInputElement txtIdentification;

  EntityInfoForm(){
    txtName = querySelector("#txtName");
    ddlProfileKey = querySelector("#ddlProfileKey");
    txtIdentification = querySelector("#txtIdentification");

    //btnaddressadd
  }

  String get name {
    return txtName.value;
  }

  Key get profileKey {
    return Key ddlProfileKey.value
  }

}