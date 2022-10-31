import { useEffect, Suspense, useState, useContext } from 'react';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
// @mui
import { styled } from '@mui/material/styles';
import { Card, Link, Container, Typography } from '@mui/material';
// hooks
import useResponsive from '../hooks/useResponsive';
// components
import Page from '../components/Page';
import Logo from '../components/Logo';
// sections
import { RegisterForm } from '../sections/auth/register';
import AuthSocial from '../sections/auth/AuthSocial';
import { LoginContext } from '../utils/Context';
// ----------------------------------------------------------------------

const RootStyle = styled('div')(({ theme }) => ({
  [theme.breakpoints.up('md')]: {
    display: 'flex',
  },
}));

const HeaderStyle = styled('header')(({ theme }) => ({
  top: 0,
  zIndex: 9,
  lineHeight: 0,
  width: '100%',
  display: 'flex',
  alignItems: 'center',
  position: 'absolute',
  padding: theme.spacing(3),
  justifyContent: 'space-between',
  [theme.breakpoints.up('md')]: {
    alignItems: 'flex-start',
    padding: theme.spacing(7, 5, 0, 7),
  },
}));

const SectionStyle = styled(Card)(({ theme }) => ({
  width: '100%',
  maxWidth: 464,
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'center',
  margin: theme.spacing(2, 0, 2, 2),
}));

const ContentStyle = styled('div')(({ theme }) => ({
  maxWidth: 480,
  margin: 'auto',
  minHeight: '100vh',
  display: 'flex',
  justifyContent: 'center',
  flexDirection: 'column',
  padding: theme.spacing(12, 0),
}));

// ----------------------------------------------------------------------

export default function Register() {
  const navigate = useNavigate();
  const { loggedIn, setLoggedIn } = useContext(LoginContext);

  const smUp = useResponsive('up', 'sm');

  const mdUp = useResponsive('up', 'md');

  const [isShown, setIsShown] = useState(false);
  useEffect(() => {
    const interval = setInterval(() => {
      if (loggedIn) {
        navigate('/auctee/home');
      } else {
        setIsShown(true);
      }
    }, 100);

    return () => clearInterval(interval);
  }, [loggedIn, isShown]);
  return (
    <>
      {isShown && (
        <Suspense fallback={<></>}>
          <Page title="Đăng ký">
            <RootStyle>
              <HeaderStyle>
                <Logo />
                {smUp && (
                  <Typography variant="body2" sx={{ mt: { md: -2 } }}>
                    Đã có tài khoản?
                    <Link variant="subtitle2" component={RouterLink} to="/auctee/login">
                      Đăng nhập
                    </Link>
                  </Typography>
                )}
              </HeaderStyle>

              {mdUp && (
                <SectionStyle>
                  <Typography variant="h3" sx={{ px: 5, mt: 10, mb: 5 }}>
                    Đấu giá và mua hàng trực tuyến với Auctee
                  </Typography>
                  <img alt="register" src="/static/illustrations/illustration_register.png" />
                </SectionStyle>
              )}

              <Container>
                <ContentStyle>
                  <Typography variant="h4" gutterBottom>
                    Đăng ký
                  </Typography>

                  <Typography sx={{ color: 'text.secondary', mb: 5 }}>Vui lòng điền thông tin cá nhân</Typography>

                  <AuthSocial />

                  <RegisterForm />

                  <Typography variant="body2" align="center" sx={{ color: 'text.secondary', mt: 3 }}>
                    By registering, I agree to Auctee &nbsp;&nbsp;
                    <Link underline="always" color="text.primary" href="#">
                      Terms of Service
                    </Link>
                    &nbsp;&nbsp;and&nbsp;&nbsp;
                    <Link underline="always" color="text.primary" href="#">
                      Privacy Policy
                    </Link>
                    .
                  </Typography>

                  {!smUp && (
                    <Typography variant="body2" sx={{ mt: 3, textAlign: 'center' }}>
                      Đã có tài khoản?
                      <Link variant="subtitle2" to="/login" component={RouterLink}>
                        Đăng nhập
                      </Link>
                    </Typography>
                  )}
                </ContentStyle>
              </Container>
            </RootStyle>
          </Page>
        </Suspense>
      )}
    </>
  );
}
