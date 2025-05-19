# rp_mgmt - Verwaltungstool für deine Rollenspielcharactere
Das Projekt rp_mgmt ist eine praktische Lösung für Rollenspiel-Enthusiasten, die ihre Charaktere effizient verwalten möchten. Es basiert auf Golang und nutzt die Fyne-GUI-Bibliothek, um eine benutzerfreundliche Oberfläche zu bieten. Die Flexibilität wird durch die Verwendung von JSON-Konfigurationsdateien gewährleistet, die es den Nutzern ermöglichen, Eigenschaften und Regeln individuell anzupassen. Dadurch können Charaktere mühelos erstellt und verwaltet werden, was das Projekt zu einer nützlichen und anpassbaren Plattform für das Rollenspiel-Management macht. :trollface:

## ToDo
* Laden von Characteren
* Speichern von Characteren
* Hochleveln von Characteren
* Export von Characteren als PDF-Dokument

## HowTo
Damit die Applikation genutzt werden kann, müssen vorab verschiedene Daten erfasst und gespeichert werden.

### Ordnerstruktur
* Linux
  * ~/.configRRoleplayManagement
  * ~/.config/RoleplayManagement/characters
  * ~/.config/RoleplayManagement/data
  * ~/.config/RoleplayManagement/logfiles
  * ~/.config/RoleplayManagement/rules
  * ~/.config/RoleplayManagement/settings
* Windows
  * %Appdata%\RoleplayManagement
  * %Appdata%\RoleplayManagement\characters
  * %Appdata%\RoleplayManagement\data
  * %Appdata%\RoleplayManagement\logfiles
  * %Appdata%\RoleplayManagement\rules
  * %Appdata%\RoleplayManagement\settings

### ./characters
Hier werden die fertig erstellten Charactere im JSON-Format gespeichert

### ./data
In diesem Ordner müssen 2 Dateien gespeichert werden
* characterproperties.json
* types.json
