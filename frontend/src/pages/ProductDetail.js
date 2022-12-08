import { useState, useEffect, lazy } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useParams, useSearchParams } from 'react-router-dom';
import { ImageViewer } from 'react-image-viewer-dv';
// material
import { Container, Rating, Typography, Stack, Button, Avatar, Link } from '@mui/material';

import { styled } from '@mui/material/styles';
import { Box } from '@mui/system';
import CountDown from '../utils/countdown';

const Page = lazy(() => import('../components/Page'));
const BidSection = lazy(() => import('../sections/bid'));

const RootStyle = styled('div')(({ theme }) => ({
  [theme.breakpoints.up('md')]: {
    display: 'flex',
    flexDirection: 'column',
  },
}));

export default function ProductDetail() {
  const navigate = useNavigate();
  const location = useLocation();

  const star = Math.floor(Math.random() * (4.9 - 3.7 + 0.5)) + 3.7;

  // get data from <server></server>
  const [searchParams, setSearchParams] = useSearchParams();
  const auctionId = searchParams.get('id');
  const productId = searchParams.get('product');

  const [ownerData, setOwnerData] = useState();
  const [auction, setAuction] = useState();
  const [product, setProduct] = useState();
  const [isFetching, setIsFetching] = useState(true);

  const handleFetchOwnerData = async (id) => {
    await fetch(`http://localhost:8080/auctee/user?id=${id && id}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setOwnerData(data);
          setIsFetching(false);
        });
      }
    });
  };

  const handleFetchAuctionData = async (id) => {
    await fetch(`http://localhost:8080/auctee/auction/detail?id=${id && id}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',

      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setAuction(data);
        });
      }
      if (res.status === 500 || res.status === 400) {
        setAuction();
        navigate('/404');
      }
    });
  };

  const handleFetchProductData = async (id) => {
    await fetch(`http://localhost:8080/auctee/product/detail?id=${id && id}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setProduct(data);
          setIsFetching(false);
        });
      }
      if (res.status === 500 || res.status === 400) {
        setProduct();
        navigate('/404');
      }
    });
  };

  useEffect(() => {
    if (isFetching) {
      handleFetchAuctionData(auctionId);
      handleFetchProductData(productId);
    }
  }, [isFetching]);

  useEffect(() => {
    if (auction && auction.user_id) {
      handleFetchOwnerData(auction.user_id);
    }
  }, [product, isFetching]);

  return isFetching ? (
    <></>
  ) : (
    <Page title={auction.name}>
      <RootStyle>
        <Container
          sx={{
            mx: '176x',
            backgroundColor: 'white',
            height: '100%',
            display: 'flex',
            minHeight: '580px',
          }}
        >
          {/* Images */}
          <Box sx={{ backgroundColor: 'transparent', flex: 2 }}>
            <Stack sx={{ py: 3, height: 'fit-content', justifyContent: 'space-between' }}>
              <Stack sx={{ height: '100%', width: '123spx', display: 'flex', alignItems: 'center', borderRadius: 0 }}>
                <ImageViewer>
                  <img
                    style={{ width: '100%', height: '100%', objectFit: 'cover', justifyContent: 'center' }}
                    src={auction.image_path}
                    alt=""
                  />
                </ImageViewer>
              </Stack>
              <Stack direction="row" sx={{ mt: 0.5 }}>
                <ImageViewer>
                  <img
                    style={{ width: '100%', height: '100%', objectFit: 'cover' }}
                    src={product.product_images.filter((item) => item.path !== auction.image_path)[0]?.path}
                    alt=""
                  />
                </ImageViewer>
                &nbsp;
                <ImageViewer>
                  <img
                    style={{ width: '100%', height: '100%', objectFit: 'cover' }}
                    src={product.product_images.filter((item) => item.path !== auction.image_path)[1]?.path}
                    alt=""
                  />
                </ImageViewer>
              </Stack>
            </Stack>
          </Box>
          {/* Detail */}
          <Box sx={{ flex: 3, ml: 3 }}>
            <Stack spacing={2} sx={{ px: 1.5, py: 2, height: '100%' }}>
              <Typography sx={{ fontSize: '1.5rem', textOverflow: 'ellipsis' }} variantx="body1 " noWrap>
                {auction.name}
              </Typography>
              {/* Rating */}
              <Stack color="#f5b727" direction="row" spacing={1}>
                <Typography sx={{ textDecoration: 'underline', fontWeight: 600 }}>{star}</Typography>
                <Rating name="half-rating-read" value={star} precision={0.1} readOnly />
              </Stack>
              {/* Price */}
              <Stack color="#f72d2d" sx={{ mt: 3 }} direction="row" alignItems="center">
                <Typography sx={{ fontSize: '1.5rem' }} variant="subtitle1">
                  {product.min_price.toLocaleString('tr-TR', {
                    style: 'currency',
                    currency: 'VND',
                  })}
                </Typography>
                <Typography sx={{ px: 1, fontSize: '1.5rem' }} variant="subtitle1">
                  -
                </Typography>
                <Typography sx={{ fontSize: '1.5rem' }} variant="subtitle1">
                  {product.expect_price.toLocaleString('tr-TR', {
                    style: 'currency',
                    currency: 'VND',
                  })}
                </Typography>
              </Stack>
              {/* Quantity */}
              <Stack direction="row" justifyContent="space-between">
                <Typography variant="body1 ">Số lượng: &nbsp; {product.quantity} </Typography>
                {new Date(auction.end_time) < new Date() ? (
                  <Typography
                    fontWeight={600}
                    fontSize={'1.2rem'}
                    variant="body1"
                    direction="row"
                    color="error"
                    sx={{
                      opacity: 0.8,
                      px: 1.5,
                      pb: 0.2,
                      border: '2px solid #f44336',
                      borderRadius: 1,
                      // maxHeight: '22.63px',
                    }}
                  >
                    Đã kết thúc
                  </Typography>
                ) : (
                  <Stack
                    // fontSize={'1rem'}
                    variant="body2"
                    direction="row"
                    sx={{
                      color: 'inherit',
                      maxHeight: '22.63px',
                    }}
                  >
                    Thời gian kết thúc :&nbsp;
                    <CountDown time={auction.end_time} />
                  </Stack>
                )}
              </Stack>
              {/* Model */}
              <Stack direction="row" alignItems="center" sx={{ maxHeight: '62px', overflow: 'hidden' }}>
                <Typography variant="body1 ">Phân loại: &nbsp;</Typography>
                <Stack direction="row">
                  {product.product_options.map((item, index) => {
                    return (
                      <Button
                        key={index}
                        sx={{
                          color: 'inherit',
                          margin: '0 5px',
                          padding: '2px 8px',
                          borderRadius: 0,
                          border: '1px solid',
                        }}
                      >
                        {item.color}
                        &nbsp; (size {item.size})
                      </Button>
                    );
                  })}
                </Stack>
              </Stack>
              {/* Bidding section */}
              <BidSection product={product} auction={auction} />
            </Stack>
          </Box>
        </Container>
        {/* second information */}
        <Container
          sx={{
            mt: 3,
            py: 2,
            mx: '176x',
            backgroundColor: 'white',
            height: '100%',
            display: 'flex',
            minHeight: '100%',
          }}
        >
          {ownerData && (
            <Stack>
              {/* Owner */}
              <Stack sx={{ mb: 2 }}>
                <Typography variant="h6">Chủ sở hữu</Typography>
                <Stack direction="row">
                  <Avatar alt="Remy Sharp" src={ownerData.avatar} sx={{ width: 40, height: 40 }} />
                  <Stack sx={{ ml: 2 }}>
                    <Typography variant="body1">{ownerData.shopname}</Typography>
                    <Link
                      sx={{ color: 'grey !important' }}
                      variant="caption"
                      component={RouterLink}
                      to="/auctee/user/profile"
                    >
                      Thông tin chi tiết
                    </Link>
                  </Stack>
                </Stack>
              </Stack>
              {/* Description */}
              {product.description && (
                <>
                  <Typography variant="h5" sx={{ mb: 1 }}>
                    Mô tả chi tiết
                  </Typography>
                  <Typography variant="body1 ">{product.description}</Typography>
                </>
              )}
            </Stack>
          )}
        </Container>
      </RootStyle>
    </Page>
  );
}
