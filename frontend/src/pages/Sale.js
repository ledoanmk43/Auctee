import { useState, useEffect, lazy } from 'react';
import { Link, useNavigate, useLocation, useSearchParams, useOutletContext } from 'react-router-dom';
import { useForm } from 'react-hook-form';
// material
import { Button, Typography, Stack, Tabs, Tab, Divider, Input } from '@mui/material';
import { LoadingButton } from '@mui/lab';
import RemoveShoppingCartOutlinedIcon from '@mui/icons-material/RemoveShoppingCartOutlined';
import StorefrontIcon from '@mui/icons-material/Storefront';
import { Icon } from '@iconify/react';
import { styled, useTheme, alpha } from '@mui/material/styles';
import { Box } from '@mui/system';
import PropTypes from 'prop-types';
import account from '../API/account';
import { FormProvider, RHFTextField } from '../components/hook-form';

import Iconify from '../components/Iconify';

const Page = lazy(() => import('../components/Page'));

const RootStyle = styled('div')(({ theme }) => ({
  [theme.breakpoints.up('md')]: {
    display: 'flex',
    flexDirection: 'column',
    mx: '176x',
    backgroundColor: 'white',
    height: '100%',
    minHeight: '580px',
  },
}));
export default function Sale() {
  const theme = useTheme();
  const navigate = useNavigate();
  const location = useLocation();

  const userData = useOutletContext();
  const [isFetching, setIsFetching] = useState(true);
  const [paymentsData, setPaymentsData] = useState();

  // Get user's data base on access_token
  const handleFetchPaymentData = async () => {
    await fetch('http://localhost:8080/auctee/user/checkout/all-bills?page=1', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setPaymentsData(data);
          setIsFetching(false);
        });
      }
      if (res.status === 401) {
        alert('You need to login first');
        setIsFetching(true);
        navigate('/auctee/login', { replace: true });
      }
    });
  };

  const handleShopCancel = async (payment) => {
    console.log(payment);
    await fetch(
      `http://localhost:8080/auctee/user/checkout/cancel-payment?id=${payment.id}&winner_id=${payment.winner_id}`,
      {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',

        mode: 'cors',
      }
    ).then((res) => {
      if (res.status === 200) {
        navigate(0);
      }

      if (res.status === 401) {
        alert('You need to login first');
        navigate('/auctee/login', { replace: true });
      }
    });
  };

  const handleConfirmDone = async (payment) => {
    await fetch(`http://localhost:8080/auctee/user/checkout/checkout-status-done?id=${payment.id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        setIsFetching(false);
        navigate(0);
      }
      if (res.status === 401) {
        alert('You need to login first');
        setIsFetching(true);
        navigate('/auctee/login', { replace: true });
      }
    });
  };
  const handleConfirmDelivery = async (payment) => {
    await fetch(`http://localhost:8080/auctee/user/checkout/shipping-confirm?id=${payment.id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',

      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        setIsFetching(false);
        navigate(0);
      }
      if (res.status === 401) {
        alert('You need to login first');
        setIsFetching(true);
        navigate('/auctee/login', { replace: true });
      }
    });
  };

  const [value, setValue] = useState(0);
  const [filteredData, setFilteredData] = useState([]);

  const handleChange = (event, newValue) => {
    setValue(newValue);
    setFilteredData([]);
    if (paymentsData) {
      switch (newValue) {
        // chờ xác nhận
        case 1:
          paymentsData.forEach((payment) => {
            if (payment.checkout_status === 1 || payment.shipping_status === 1) {
              setFilteredData((current) => [...current, payment]);
            }
          });
          break;
        // đang giao
        case 2:
          paymentsData.forEach((payment) => {
            if (payment.checkout_status === 3 && payment.shipping_status === 2) {
              setFilteredData((current) => [...current, payment]);
            }
          });
          break;
        // đã giao
        case 3:
          paymentsData.forEach((payment) => {
            if ((payment.checkout_status === 3 && payment.shipping_status === 3) || payment.checkout_status === 5) {
              setFilteredData((current) => [...current, payment]);
            }
          });
          break;
        // đã huỷ
        case 4:
          paymentsData.forEach((payment) => {
            if (payment.checkout_status === 4) {
              setFilteredData((current) => [...current, payment]);
            }
          });
          break;
        case 5:
          paymentsData.forEach((payment) => {
            if (payment.checkout_status === 4) {
              setFilteredData((current) => [...current, payment]);
            }
          });
          break;
        default:
          setFilteredData(paymentsData);
          break;
      }
    }
  };

  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    !paymentsData && handleFetchPaymentData();
    setFilteredData(paymentsData);
  }, [isFetching, paymentsData]);

  return !isFetching ? (
    <Page title="Bán hàng">
      <RootStyle sx={{ px: 3, py: 2, maxWidth: '980px', bgcolor: 'red' }}>
        {/* Heading */}
        <Stack>
          <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
            Đơn hàng và Vận chuyển
          </Typography>
          <Typography variant="body2" sx={{ color: 'black', opacity: 0.8, pb: 2 }}>
            Quản lý các đơn hàng của bạn một cách thuận tiện nhất
          </Typography>
          <Divider />
        </Stack>
        {/* Main */}
        <Stack sx={{ pt: 2 }} direction="row" alignItems="center" justifyContent="space-between">
          <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
            {paymentsData?.length} Đơn hàng
          </Typography>
          <SearchbarStyle>
            <Input
              disableUnderline
              placeholder="Tìm kiếm đơn hàng"
              inputProps={{
                sx: {
                  '&::placeholder': {
                    fontSize: '0.87rem',
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
                mr: -4.2,
                py: 0.75,
                px: 2,
                backgroundColor: `${alpha(theme.palette.background.main, 0.9)}`,
              }}
            >
              <Iconify icon="eva:search-fill" sx={{ color: 'white', width: 20, height: 20, fontSize: '0.9rem' }} />
            </Button>
          </SearchbarStyle>
        </Stack>
        <Stack sx={{ mt: 1 }}>
          <Box sx={{ width: '100%' }}>
            <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
              <Tabs
                sx={{
                  '& button': {
                    minWidth: 146,
                    color: 'inherit',
                    fontSize: '0.9rem',
                    fontWeight: 400,
                    textTransform: 'none',
                    px: 4,
                  },
                  '& button:hover': {
                    color: '#f44336',
                  },
                  '& button.Mui-selected': {
                    color: '#f44336',
                  },
                }}
                textColor="primary"
                TabIndicatorProps={{ style: { background: '#f44336' } }}
                value={value}
                onChange={handleChange}
                aria-label="secondary tabs example"
              >
                <Tab disableRipple label="Tất cả" {...a11yProps(0)} />
                <Tab disableRipple label="Chờ xác nhận" {...a11yProps(1)} />
                <Tab disableRipple label="Đang giao" {...a11yProps(2)} />
                <Tab disableRipple label="Đã giao" {...a11yProps(3)} />
                <Tab disableRipple label="Đơn huỷ" {...a11yProps(4)} />
                <Tab disableRipple label="Trả hàng/Hoàn tiền" {...a11yProps(5)} />
              </Tabs>
            </Box>
            <TabPanel value={value} index={value}>
              {filteredData ? (
                filteredData.map((payment, index) => (
                  <Stack sx={{ boxShadow: 4, mb: 2, p: 2 }} direction="column" key={index}>
                    {/* Top side */}
                    <Stack maxHeight={20} sx={{ mb: 0.5 }} direction="row" justifyContent="space-between">
                      <Stack direction="row">
                        <Typography
                          variant="button"
                          sx={{
                            textTransform: 'none',
                            bgcolor: '#f44336',
                            color: 'white',
                            borderRadius: 0.5,
                            fontSize: '0.7rem',
                            px: 0.5,
                            mr: 1.5,
                          }}
                        >
                          Shop yêu thích
                        </Typography>
                        <Link
                          style={{
                            fontWeight: 600,
                            color: 'inherit',
                            textDecoration: 'none',
                          }}
                        >
                          {payment.shopname}
                        </Link>
                        <Button
                          disableRipple
                          sx={{ ml: 1.5, border: '1px solid black', borderRadius: 0.4, opacity: 0.9, color: 'inherit' }}
                        >
                          <StorefrontIcon sx={{ fontSize: '1rem' }} />
                          <Typography
                            variant="button"
                            sx={{
                              textTransform: 'none',
                              fontSize: '0.7rem',
                              px: 0.5,
                            }}
                          >
                            Shop của tôi
                          </Typography>
                        </Button>
                        {payment.payment_method.length > 0 && (
                          <Button
                            disableRipple
                            sx={{
                              ml: 1.5,
                              borderRadius: 0.4,
                              opacity: 0.9,
                              textTransform: 'none',
                              fontSize: '0.75rem',
                              fontStyle: 'italic',
                              color: `${payment.payment_method.length !== 3 ? '#a50064' : 'red'}`,
                              '&:hover': {
                                bgcolor: 'transparent',
                              },
                            }}
                          >
                            {payment.payment_method}
                          </Button>
                        )}
                        {payment.shipping_status === 3 && payment.total !== 0 && (
                          <Button
                            disableRipple
                            sx={{
                              ml: 1.5,
                              borderRadius: 0.4,
                              textTransform: 'none',
                              fontSize: '0.75rem',
                              fontStyle: 'italic',
                              opacity: 0.7,
                              '&:hover': {
                                bgcolor: 'transparent',
                              },
                            }}
                          >
                            Đã nhận hàng
                          </Button>
                        )}
                        {payment.checkout_status === 4 && (
                          <Button
                            variant="error"
                            disableRipple
                            sx={{
                              px: 0,
                              ml: 1.5,
                              borderRadius: 0.4,
                              opacity: 0.9,
                              color: 'inherit',
                              '&:hover': {
                                bgcolor: 'transparent',
                              },
                            }}
                          >
                            <RemoveShoppingCartOutlinedIcon color="error" sx={{ fontSize: '1rem' }} />
                            <Typography
                              color="error"
                              variant="button"
                              sx={{
                                textTransform: 'none',
                                fontSize: '0.7rem',
                                px: 0.5,
                              }}
                            >
                              Đã huỷ
                            </Typography>
                          </Button>
                        )}
                      </Stack>
                      <Stack justifyContent="flex-end" direction="row" alignItems="flex-end">
                        <Typography sx={{ fontSize: '0.85rem' }} variant="caption">
                          ID: &nbsp;
                        </Typography>
                        <Typography sx={{ fontSize: '0.85rem' }} variant="caption">
                          {payment.id}
                        </Typography>
                      </Stack>
                    </Stack>
                    {/* Body */}
                    <Stack justifyContent="space-between" direction="row" sx={{ display: 'flex' }}>
                      {/* Image */}
                      <Stack flex={4} direction="row">
                        <Stack>
                          <ProductImgStyle alt={payment.product_id} src={payment.image_path} />
                        </Stack>
                        {/* Name */}
                        <Stack sx={{ mx: 2 }}>
                          <Typography variant="caption" sx={{ textTransform: 'uppercase', fontSize: '1rem' }}>
                            {payment.product_name}
                          </Typography>
                          <Typography sx={{ mt: 1 }}>x{payment.quantity}</Typography>
                        </Stack>
                      </Stack>
                      {/* Total */}
                      <Stack alignItems="flex-end" flex={1.7}>
                        <Typography sx={{ fontSize: '0.85rem' }} variant="caption">
                          Tổng số tiền tạm tính:
                        </Typography>
                        <Typography color="#f44336">
                          {(payment.before_discount + payment.shipping_value).toLocaleString('tr-TR', {
                            style: 'currency',
                            currency: 'VND',
                          })}
                        </Typography>
                        {payment.shipping_status === 1 && (
                          <Stack sx={{}} width="100%" direction="row" justifyContent="flex-end">
                            <Button
                              color="error"
                              size="medium"
                              variant="contained"
                              disableRipple
                              sx={{
                                borderRadius: 0.4,
                                bgcolor: '#f44336',
                                color: 'white',
                                px: 1.5,
                                textTransform: 'none',
                              }}
                              onClick={() => handleConfirmDelivery(payment)}
                            >
                              Giao cho đơn vị vận chuyển
                            </Button>
                          </Stack>
                        )}
                        {payment.shipping_status === 2 && (
                          <Stack sx={{}} width="100%" direction="row" justifyContent="flex-end">
                            <Button
                              color="error"
                              size="medium"
                              variant="contained"
                              disableRipple
                              sx={{
                                borderRadius: 0.4,
                                bgcolor: '#f44336',
                                color: 'white',
                                px: 1.5,
                                textTransform: 'none',
                              }}
                              onClick={() => handleShopCancel(payment)}
                            >
                              Không giao được hàng
                            </Button>
                          </Stack>
                        )}
                        {((payment.checkout_status === 3 && payment.shipping_status === 3) ||
                          payment.checkout_status === 5) && (
                          <Stack sx={{}} width="100%" direction="row" justifyContent="flex-end">
                            <Button
                              disabled={payment.checkout_status === 5}
                              color="error"
                              size="medium"
                              variant="contained"
                              disableRipple
                              sx={{
                                borderRadius: 0.4,
                                bgcolor: '#f44336',
                                color: 'white',
                                px: 1.5,
                                textTransform: 'none',
                                '&:disabled': {
                                  backgroundColor: `${payment.checkout_status === 5 && '  '}`,
                                  color: `${payment.checkout_status === 5 && 'white'}`,
                                  opacity: `${payment.checkout_status === 5 && '0.8'}`,
                                },
                              }}
                              onClick={() => handleConfirmDone(payment)}
                            >
                              {payment.checkout_status === 5 ? 'Đã hoàn thành' : 'Xác nhận hoàn thành đơn hàng'}
                            </Button>
                          </Stack>
                        )}
                      </Stack>
                    </Stack>
                  </Stack>
                ))
              ) : (
                <>Chưa có thông tin</>
              )}
            </TabPanel>
          </Box>
        </Stack>
      </RootStyle>
    </Page>
  ) : (
    <>Có lỗi xảy ra</>
  );
}
function TabPanel(props) {
  const { children, value, index, ...other } = props;
  const [statusText, setStatusText] = useState('');
  useEffect(() => {
    switch (value) {
      case 1:
        setStatusText('Có vẻ bạn chưa có đơn hàng nào đang chờ xác nhận');
        break;
      case 2:
        setStatusText('Có vẻ bạn chưa có đơn hàng nào đang giao');
        break;
      case 3:
        setStatusText('Không tìm thấy dữ liệu đơn hàng đã giao');
        break;
      case 4:
        setStatusText('Không tìm thấy đơn đã huỷ');
        break;

      default:
        setStatusText('Có vẻ bạn chưa có đơn hàng nào');

        break;
    }
  }, [value]);
  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box sx={{ py: 3 }}>
          <Stack>
            {children.length > 0 ? (
              children
            ) : (
              <Stack fontStyle="italic" sx={{ mx: 'auto' }}>
                {statusText}
              </Stack>
            )}
          </Stack>
        </Box>
      )}
    </div>
  );
}

TabPanel.propTypes = {
  children: PropTypes.node,
  index: PropTypes.number.isRequired,
  value: PropTypes.number.isRequired,
};

function a11yProps(index) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}
const ProductImgStyle = styled('img')({
  width: '85px',
  height: '85px',
  objectFit: 'cover',
});
const APPBAR_MOBILE = 32;
const APPBAR_DESKTOP = 46;
const DRAWER_WIDTH = 660;

const SearchbarStyle = styled('div')(({ theme }) => ({
  width: 250,
  display: 'flex',

  border: '1px solid grey',
  justifyContent: 'center',
  alignItems: 'center',
  height: APPBAR_MOBILE,
  padding: theme.spacing(0, 4),
  backgroundColor: 'white',
}));
