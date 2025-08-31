import 'package:flutter/material.dart';

final darkThemeData = ThemeData(
  brightness: Brightness.dark,
  primarySwatch: Colors.deepPurple,
  appBarTheme: AppBarTheme(
    color: Colors.deepPurple[700],
    elevation: 2,
  ),
  cardTheme: CardTheme(
    elevation: 4,
    shape: RoundedRectangleBorder(
      borderRadius: BorderRadius.circular(10),
    ),
  ),
);
