import React, { createContext, useState, useContext, useEffect, type ReactNode } from 'react';

interface ThemeColors {
  background: string;
  onBackground: string;

  primary: string;
  onPrimary: string;
  primaryContainer: string;
  onPrimaryContainer: string;

  secondary: string;
  onSecondary: string;
  secondaryContainer: string;
  onSecondaryContainer: string;

  tertiary: string;
  onTertiary: string;
  tertiaryContainer: string;
  onTertiaryContainer: string;

  error: string;
  onError: string;
  errorContainer: string;
  onErrorContainer: string;

  outline: string;
}

interface ThemeContextType {
  theme: 'light' | 'dark';
  toggleTheme: () => void;
  colors: ThemeColors;
}

const lightThemeColors = {
  background: '#f6fbf4',
  onBackground: '#171d19',

  primary: '#296a47',
  onPrimary: '#ffffff',
  primaryContainer: '#aef2c6',
  onPrimaryContainer: '#085131',

  secondary: '#4e6355',
  onSecondary: '#ffffff',
  secondaryContainer: '#d1e8d6',
  onSecondaryContainer: '#374b3e',

  tertiary: '#3b6471',
  onTertiary: '#ffffff',
  tertiaryContainer: '#bfe9f8',
  onTertiaryContainer: '#224c58',

  error: '#ba1a1a',
  onError: '#ffffff',
  errorContainer: '#ffdad6',
  onErrorContainer: '#93000a',

  outline: '#717972',
};

const darkThemeColors = {
  background: '#0f1511',
  onBackground: '#dfe4dd',

  primary: '#92d5ab',
  onPrimary: '#003920',
  primaryContainer: '#085131',
  onPrimaryContainer: '#aef2c6',

  secondary: '#b5ccba',
  onSecondary: '#213528',
  secondaryContainer: '#374b3e',
  onSecondaryContainer: '#d1e8d6',

  tertiary: '#a3cddc',
  onTertiary: '#033541',
  tertiaryContainer: '#224c58',
  onTertiaryContainer: '#bfe9f8',

  error: '#ffb4ab',
  onError: '#690005',
  errorContainer: '#93000a',
  onErrorContainer: '#ffdad6',

  outline: '#8a938b',
};

const ThemeContext = createContext<ThemeContextType>({
  theme: 'light',
  toggleTheme: () => {},
  colors: lightThemeColors,
});

interface ThemeProviderProps {
  children: ReactNode;
}

export const ThemeProvider: React.FC<ThemeProviderProps> = ({ children }) => {
  const [theme, setTheme] = useState<'light' | 'dark'>(() => {
    const savedTheme = localStorage.getItem('theme');
    return (savedTheme === 'light' || savedTheme === 'dark') ? savedTheme : 'light';
  });

  useEffect(() => localStorage.setItem('theme', theme), [theme]);

  const toggleTheme = () => {
    setTheme((prevTheme) => (prevTheme === 'light' ? 'dark' : 'light'));
  };

  const colors = theme === 'light' ? lightThemeColors : darkThemeColors;
  document.documentElement.style.setProperty('--background-color', colors.background);
  document.documentElement.style.setProperty('--on-background-color', colors.onBackground);
  document.documentElement.style.setProperty('--primary-color', colors.primary);
  document.documentElement.style.setProperty('--on-primary-color', colors.onPrimary);
  document.documentElement.style.setProperty('--primary-container-color', colors.primaryContainer);
  document.documentElement.style.setProperty('--on-primary-container-color', colors.onPrimaryContainer);
  document.documentElement.style.setProperty('--secondary-color', colors.secondary);
  document.documentElement.style.setProperty('--on-secondary-color', colors.onSecondary);
  document.documentElement.style.setProperty('--secondary-container-color', colors.secondaryContainer);
  document.documentElement.style.setProperty('--on-secondary-container-color', colors.onSecondaryContainer);
  document.documentElement.style.setProperty('--tertiary-color', colors.tertiary);
  document.documentElement.style.setProperty('--on-tertiary-color', colors.onTertiary);
  document.documentElement.style.setProperty('--tertiary-container-color', colors.tertiaryContainer);
  document.documentElement.style.setProperty('--on-tertiary-container-color', colors.onTertiaryContainer);
  document.documentElement.style.setProperty('--error-color', colors.error);
  document.documentElement.style.setProperty('--on-error-color', colors.onError);
  document.documentElement.style.setProperty('--error-container-color', colors.errorContainer);
  document.documentElement.style.setProperty('--on-error-container-color', colors.onErrorContainer);
  document.documentElement.style.setProperty('--outline-color', colors.outline);


  const contextValue: ThemeContextType = {
    theme,
    toggleTheme,
    colors,
  };

  return (
    <ThemeContext.Provider value={contextValue}>
      {children}
    </ThemeContext.Provider>
  );
};

export const useTheme = (): ThemeContextType => {
  return useContext(ThemeContext);
};
