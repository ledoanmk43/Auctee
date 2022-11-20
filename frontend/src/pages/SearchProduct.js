import { useState, useEffect, lazy, Suspense } from 'react';
import { Link as RouterLink, useNavigate, useSearchParams, useLocation } from 'react-router-dom';

// material
import { Box, Button, Link, Typography, Stack, Container } from '@mui/material';
import { styled } from '@mui/material/styles';
import { ProductSort, ProductList, ProductCartWidget, ProductFilterSidebar } from '../sections/@dashboard/products';

const Page = lazy(() => import('../components/Page'));

const ContentStyle = styled('div')(({ theme }) => ({
  maxWidth: 600,
  margin: 'auto',
  minHeight: '80vh',
  display: 'flex',
  justifyContent: 'center',
  flexDirection: 'column',
  padding: theme.spacing(0, 0),
}));

export default function SearchProduct() {
  const navigate = useNavigate();
  const location = useLocation();

  const [openFilter, setOpenFilter] = useState(false);
  const handleOpenFilter = () => {
    setOpenFilter(true);
  };

  const handleCloseFilter = () => {
    setOpenFilter(false);
  };

  const [auctionsData, setAuctionsData] = useState([]);

  const [searchParams, setSearchParams] = useSearchParams();
  const kwd = searchParams.get('keyword');
  const [isFetching, setIsFetching] = useState(true);

  const handleSearchProduct = async (kw) => {
    await fetch(`http://localhost:1009/auctee/auctions/products?product_name=${kw.toLowerCase()}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setAuctionsData(data);
          setIsFetching(false);
        });
      }
      if (res.status === 404) {
        setAuctionsData();
        setIsFetching(true);
      }
      if (res.status === 400) {
        setIsFetching(true);
        setAuctionsData();
      }
    });
  };

  useEffect(() => {
    if (kwd === '' || kwd === null) {
      navigate('/auctee/home', { replace: true });
    }
    if (isFetching || location.pathname === '/auctee/search/') {
      handleSearchProduct(kwd);
    }
  }, [kwd, searchParams]);
  // console.log(kwd);

  return isFetching ? (
    <Suspense startTransition callback={<></>}>
      <Page title="Search">
        <Container>
          <ContentStyle sx={{ textAlign: 'center', alignItems: 'center' }}>
            <Typography variant="h3" paragraph>
              Xin lỗi, không tìm thấy sản phẩm!
            </Typography>

            <Box
              component="img"
              src="/static/illustrations/illustration_404.svg"
              sx={{ height: 260, mx: 'auto', my: { xs: 5, sm: 10 } }}
            />

            <Button
              to="/auctee/home"
              size="large"
              sx={{
                bgcolor: '#f44336',
                '&:hover': {
                  opacity: 0.9,
                  bgcolor: 'transparent',
                  color: '#f44336',
                  border: '1px solid #f44336',
                },
              }}
              variant="contained"
              component={RouterLink}
            >
              Go to Home
            </Button>
          </ContentStyle>
        </Container>
      </Page>
    </Suspense>
  ) : (
    <Page title="Search">
      <Container>
        <Typography variant="h4" sx={{ mb: 2 }}>
          Results
        </Typography>

        <Stack direction="row" flexWrap="wrap-reverse" alignItems="center" justifyContent="flex-end" sx={{ mb: 5 }}>
          <Stack direction="row" spacing={1} flexShrink={0} sx={{ my: 1 }}>
            <ProductFilterSidebar
              isOpenFilter={openFilter}
              onOpenFilter={handleOpenFilter}
              onCloseFilter={handleCloseFilter}
            />
            <ProductSort />
          </Stack>
        </Stack>

        <ProductList auctions={auctionsData} />
      </Container>
    </Page>
  );
}
