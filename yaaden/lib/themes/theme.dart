import 'package:flutter/material.dart';
import './light.dart';
import './dark.dart';

// Theme Provider for managing app themes
class ThemeProvider extends ChangeNotifier {
  ThemeMode _themeMode = ThemeMode.system;

  ThemeMode get themeMode => _themeMode;

  void setThemeMode(ThemeMode mode) {
    _themeMode = mode;
    notifyListeners();
  }
}

// Custom App Themes
class AppThemes {
  static final darkTheme = darkThemeData;
  static final lightTheme = lightThemeData;
}
