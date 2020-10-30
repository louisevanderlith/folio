class Contact {
  final String icon;
  final String name;
  final String value;

  Contact(this.icon, this.name, this.value);

  Map<String, dynamic> toJson() {
    return {"Icon": icon, "Name": name, "Value": value};
  }
}
