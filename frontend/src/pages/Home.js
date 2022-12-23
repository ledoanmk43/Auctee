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
  const [openFilter, setOpenFilter] = useState(false);
  const listInnerRef = useRef();
  const handleOpenFilter = () => {
    setOpenFilter(true);
  };

  const handleCloseFilter = () => {
    setOpenFilter(false);
  };

  const [isFetching, setIsFetching] = useState(true);

  const [currPage, setCurrPage] = useState(1); // storing current page number
  const [prevPage, setPrevPage] = useState(0); // storing prev page number
  const [wasLastList, setWasLastList] = useState(false); // setting a flag to know the last list
  const [auctionsData, setAuctionsData] = useState([]);

  const onScroll = () => {
    if (listInnerRef.current) {
      const { scrollTop, scrollHeight, clientHeight } = listInnerRef.current;
      if (scrollTop + clientHeight === scrollHeight) {
        setCurrPage(currPage + 1);
      }
    }
  };
  // useEffect(() => {
  //   if (isFetching) {
  //     handleFetchAuctionData();
  //   }
  // }, []);

  useEffect(() => {
    const fetchData = async () => {
      // eslint-disable-next-line no-unneeded-ternary
      await fetch(`http://localhost:8080/auctee/auctions?page=${currPage ? currPage : 1}`, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        mode: 'cors',
      }).then((response) => {
        response.json().then((data) => {
          if (!data.length) {
            setWasLastList(true);
            return;
          }
          setPrevPage(currPage);
          setAuctionsData([...auctionsData, ...data]);
        });
      });
    };
    if (!wasLastList && prevPage !== currPage) {
      fetchData();
    }
  }, [currPage, wasLastList, prevPage, auctionsData]);

  return (
    <Page
      onScroll={onScroll}
      ref={listInnerRef}
      sx={{
        height: '100vh',
        overflowY: 'auto',
        '&::-webkit-scrollbar': {
          display: 'none',
        },
      }}
      title="Auctee"
    >
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
        {auctionsData && <ProductList auctions={auctionsData} />}
      </Container>
    </Page>
  );
}
