import { useState, useEffect, lazy, useContext, useRef } from 'react';
import { useLocation } from 'react-router-dom';

// material
import { Container, Stack, Typography, Button, CardMedia, Card } from '@mui/material';
// components
import { ProductSort, ProductList, ProductFilterSidebar } from '../sections/@dashboard/products';
import { LoginContext } from '../utils/Context';

const Page = lazy(() => import('../components/Page'));
// ----------------------------------------------------------------------

export default function EcommerceShop() {
  const { loggedIn, setLoggedIn } = useContext(LoginContext);
  const location = useLocation();
  // const [openFilter, setOpenFilter] = useState(false);

  // const handleOpenFilter = () => {
  //   setOpenFilter(true);
  // };

  // const handleCloseFilter = () => {
  //   setOpenFilter(false);
  // };

  const [currPage, setCurrentPage] = useState(1);

  const [loadDone, setLoadDone] = useState(false);
  const [auctionsData, setAuctionsData] = useState([]);
  const fetchData = async () => {
    // eslint-disable-next-line no-unneeded-ternary
    await fetch(`http://localhost:8080/auctee/auctions?page=${currPage}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      mode: 'cors',
    }).then((response) => {
      response.json().then((data) => {
        if (data.length === 0) {
          setLoadDone(true);
        }
        setAuctionsData((datas) => [...datas, ...data]);
      });
    });
  };
  const handleLoadMore = () => {
    setCurrentPage(currPage + 1);
  };
  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    currPage > 1 && fetchData();
  }, [currPage]);
  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    auctionsData.length === 0 && fetchData();
  }, [auctionsData]);

  return (
    <Page title="Auctee">
      <Container>
        <Stack direction="row" alignItems="center" justifyContent="center" sx={{ mb: 2 }}>
          <CardMedia component="img" height="300" image="/static/banner1.jpg" alt="banner" />

          {/* <ProductFilterSidebar
              isOpenFilter={openFilter}
              onOpenFilter={handleOpenFilter}
              onCloseFilter={handleCloseFilter}
            />
            <ProductSort /> */}
        </Stack>
        <Typography variant="h4" sx={{ mb: 2 }}>
          Tất cả sản phẩm đang được đấu giá
        </Typography>
        <Stack>
          {auctionsData && <ProductList auctions={auctionsData} />}
          {auctionsData.length !== 0 && (
            <Button
              color="error"
              onClick={() => handleLoadMore()}
              disableRipple
              sx={{ mx: 'auto', my: 2, textTransform: 'none' }}
            >
              {loadDone ? 'Đã tải xong' : 'Xem thêm'}
            </Button>
          )}
        </Stack>
      </Container>
    </Page>
  );
}
