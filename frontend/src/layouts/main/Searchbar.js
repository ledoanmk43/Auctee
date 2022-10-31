import { useState, useEffect, Suspense } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams } from 'react-router-dom';
// material
import { styled, alpha, useTheme } from '@mui/material/styles';
import { Input, Button, Stack, Typography, ClickAwayListener } from '@mui/material';
// component
import Iconify from '../../components/Iconify';
import Scrollbar from '../../components/Scrollbar';
// ----------------------------------------------------------------------

const APPBAR_MOBILE = 32;
const APPBAR_DESKTOP = 46;
const DRAWER_WIDTH = 280;

const SearchbarStyle = styled('div')(({ theme }) => ({
  top: `calc(100% - ${APPBAR_MOBILE + 9}px)`,
  zIndex: 999,
  width: `calc(100% - 2*${DRAWER_WIDTH + 41}px)`,
  display: 'flex',
  position: 'absolute',
  justifyContent: 'center',
  alignItems: 'center',
  height: APPBAR_MOBILE,
  padding: theme.spacing(0, 3),
  boxShadow: theme.customShadows.z8,
  backgroundColor: 'white',
  [theme.breakpoints.up('md')]: {
    top: `calc(100% - ${APPBAR_DESKTOP + 23}px)`,
    height: APPBAR_DESKTOP,
    padding: theme.spacing(0, 5),
  },
}));

// ----------------------------------------------------------------------

export default function Searchbar() {
  const theme = useTheme();
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();
  const kwd = searchParams.get('keyword');

  const [open, setOpen] = useState(false);
  const [keyword, setKeyWord] = useState('');
  const [recent, setRecent] = useState([]);
  const handleSearch = (e) => {
    e?.preventDefault();
    if (keyword?.length > 0) {
      /* eslint-disable-next-line no-plusplus */
      for (let i = 0; i < recent.length; i++) {
        if (recent[i] === keyword) {
          recent.splice(i, 1);
        }
        if (recent.length >= 15) {
          recent.pop();
        }
      }
      recent.unshift(keyword);
      window.localStorage.setItem('recent', JSON.stringify(recent));
      navigate(`/auctee/search/?keyword=${keyword}`);
    } else {
      navigate('/auctee/home', { replace: true });
    }
  };

  useEffect(() => {
    setKeyWord(kwd);
    setRecent(JSON.parse(window.localStorage.getItem('recent')));
  }, [kwd, open]);

  return (
    <ClickAwayListener onClickAway={() => setOpen(false)}>
      <form onSubmit={(e) => handleSearch(e)}>
        <SearchbarStyle>
          <Input
            onClick={() => setOpen(true)}
            value={keyword || ''}
            onChange={(e) => setKeyWord(e.target.value)}
            fullWidth
            onFocus={() => setOpen(true)}
            disableUnderline
            placeholder="Nhanh tay ẵm ngay giá tốt"
            inputProps={{
              sx: {
                '&::placeholder': {
                  opacity: 0.62,
                  color: 'black',
                  fontWeight: 200,
                },
              },
            }}
            sx={{
              mr: 1,
              ml: -3,
            }}
          />
          <Button
            type="submit"
            sx={{
              ':hover': {
                bgcolor: `${alpha(theme.palette.background.main, 0.8)}`,
              },
              borderRadius: 0,
              mr: -4.4,
              py: 1,
              px: 4,
              backgroundColor: `${alpha(theme.palette.background.main, 0.9)}`,
            }}
          >
            <Iconify icon="eva:search-fill" sx={{ color: 'white', width: 20, height: 20 }} />
          </Button>
          {open && (
            <Stack
              sx={{
                overflow: 'auto',
                scrollbarWidth: 'thin',
                '&::-webkit-scrollbar': {
                  width: '0.4em',
                },
                '&::-webkit-scrollbar-track': {
                  background: '#f0e7e6',
                },
                '&::-webkit-scrollbar-thumb': {
                  backgroundColor: '#cfc9c8',
                },
                '&::-webkit-scrollbar-thumb:hover': {
                  background: '#bab3b1',
                },
                maxHeight: '250px',
                width: '100%',
                minHeight: '50px',
                top: `calc(100% - ${APPBAR_MOBILE - 34}px)`,
                position: 'absolute',
                bgcolor: 'white',
              }}
            >
              {recent.map((item, index) => {
                console.log(item);
                return (
                  <Button
                    type="submit"
                    onClick={() => {
                      setKeyWord(item);
                      handleSearch();
                    }}
                    sx={{
                      justifyContent: 'flex-start',
                      color: 'black',
                      px: 2,
                      borderRadius: '0',
                      fontWeight: 500,
                      opacity: 0.7,
                      '&:hover': {
                        opacity: 0.9,
                      },
                    }}
                    key={index}
                  >
                    {item}
                  </Button>
                );
              })}
            </Stack>
          )}
        </SearchbarStyle>
      </form>
    </ClickAwayListener>
  );
}
