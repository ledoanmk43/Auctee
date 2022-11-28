import { useState, useEffect, lazy, useContext } from 'react';
import { useLocation } from 'react-router-dom';

// material
import { Container, Stack, Typography } from '@mui/material';
// components
import { ProductSort, ProductList, ProductFilterSidebar } from '../sections/@dashboard/products';
import { LoginContext } from '../utils/Context';

const Page = lazy(() => import('../components/Page'));
// ----------------------------------------------------------------------

export default function EcommerceShop() {
  const { loggedIn, setLoggedIn } = useContext(LoginContext);
  const location = useLocation();
  const [openFilter, setOpenFilter] = useState(false);

  const handleOpenFilter = () => {
    setOpenFilter(true);
  };

  const handleCloseFilter = () => {
    setOpenFilter(false);
  };

  const [isFetching, setIsFetching] = useState(true);

  const [auctionsData, setAuctionsData] = useState([]);

  const handleFetchAuctionData = async () => {
    await fetch('http://localhost:1009/auctee/auctions?page=1', {
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
      if (res.status === 500 || res.status === 400) {
        setAuctionsData();
      }
    });
  };
  useEffect(() => {
    if (isFetching) {
      handleFetchAuctionData();
    }
  }, []);

  return isFetching ? (
    <></>
  ) : (
    <Page title="Auctee">
      <Container>
        <Typography variant="h4" sx={{ mb: 2 }}>
         Tất cả sản phẩm đang được đấu giá
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
