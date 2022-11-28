// routes
import { useNavigate, useLocation } from 'react-router-dom';
import { useState, useEffect, useMemo } from 'react';
import { viVN } from '@mui/material/locale';
import * as locales from '@mui/material/locale';
import { createTheme, useTheme } from '@mui/material/styles';
import Router from './routes';
import { LoginContext, ReloadContext, CurrentAuctionsContext } from './utils/Context';
// theme
import ThemeProvider from './theme';
// components
import ScrollToTop from './components/ScrollToTop';
import { BaseOptionChartStyle } from './components/chart/BaseOptionChart';

// ----------------------------------------------------------------------

export default function App() {
  const navigate = useNavigate();
  const location = useLocation();
  const [loggedIn, setLoggedIn] = useState(false);
  const [isReloading, setIsReloading] = useState(false);

  const theme = useTheme();
  const [locale, setLocale] = useState('viVN');
  const themeWithLocale = useMemo(() => createTheme(theme, locales[locale]), [locale, theme]);

  const handleCheckLogin = async () => {
    await fetch('http://localhost:1001/auctee/refreshToken', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    }).then((res) => {
      if (res.status === 401) {
        setLoggedIn(false);
        if (location.pathname === '/auctee/register') {
          navigate('auctee/register');
        } else {
          navigate('auctee/login');
        }
      }
      if (res.status === 200) {
        setLoggedIn(true);
      }
    });
  };

  useEffect(() => {
    handleCheckLogin();
  }, []);
  return (
    <ThemeProvider theme={themeWithLocale}>
      <ScrollToTop />
      <BaseOptionChartStyle />
      <LoginContext.Provider value={{ loggedIn, setLoggedIn }}>
        <ReloadContext.Provider value={{ isReloading, setIsReloading }}>
          <Router />
        </ReloadContext.Provider>
      </LoginContext.Provider>
    </ThemeProvider>
  );
}
