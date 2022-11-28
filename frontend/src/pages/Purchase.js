import { useState, useEffect, lazy } from 'react';
import { Link, useNavigate, useLocation, useSearchParams, useOutletContext } from 'react-router-dom';
import { useForm } from 'react-hook-form';
// material
import { Button, Typography, Stack, Tabs, Tab, Divider } from '@mui/material';
import { LoadingButton } from '@mui/lab';
import StorefrontIcon from '@mui/icons-material/Storefront';
import { Icon } from '@iconify/react';
import { styled, useTheme } from '@mui/material/styles';
import { Box } from '@mui/system';
import PropTypes from 'prop-types';
import account from '../API/account';
import { FormProvider, RHFTextField } from '../components/hook-form';

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
export default function Purchase() {
  const navigate = useNavigate();
  const location = useLocation();

  const userData = useOutletContext();
  const [isFetching, setIsFetching] = useState(true);
  const [paymentsData, setPaymentsData] = useState();
  const [shopName, setShopName] = useState('');
  const [nickName, setNickName] = useState('');
  const [phoneNumber, setPhoneNumber] = useState('');
  const [avatarFile, setAvatarFile] = useState();
  const [HonorPoint, setHonorPoint] = useState(0);

  const [isMale, setIsMale] = useState(false); // 1 male : 0 female

  // Get user's data base on access_token
  const handleFetchPaymentData = async () => {
    await fetch('http://localhost:1003/auctee/user/checkout/payment-history?page=1', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          const today = new Date();
          setPaymentsData(data);
          setShopName(data.shopname);
          setNickName(data.nickname);
          setIsMale(data.gender);
          setPhoneNumber(data.phone);
          setAvatarFile(data.avatar);
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

  const defaultValues = {
    nickname: '',
    shopname: '',
    gender: isMale,
    phone: '',
    date: '',
    month: '',
    year: '',
    avatar: '',
  };

  const methods = useForm({
    defaultValues,
  });
  const {
    handleSubmit,
    formState: { isSubmitting },
  } = methods;

  const onSubmit = async () => {
    const payload = {};

    // await fetch('http://localhost:1001/auctee/user/profile/setting', {
    //   method: 'PUT',
    //   headers: { 'Content-Type': 'application/json' },
    //   body: JSON.stringify(payload),
    //   credentials: 'include',
    // }).then((res) => {
    //   if (res.status === 200) {
    //     setIsUpdated(true);
    //     navigate(0);
    //   }
    //   if (res.status === 409) {
    //     setIsUpdated(false);
    //   }
    //   if (res.status === 400) {
    //     setIsUpdated(false);
    //   }
    // });
  };
  const [value, setValue] = useState(0);
  const [filteredData, setFilteredData] = useState([]);

  const handleChange = (event, newValue) => {
    setValue(newValue);
    setFilteredData([]);
    if (paymentsData) {
      switch (newValue) {
        case 1:
          paymentsData.forEach((payment) => {
            if (payment.checkout_status === 1) {
              setFilteredData((current) => [...current, payment]);
            }
          });
          break;
        case 2:
          paymentsData.forEach((payment) => {
            if (payment.checkout_status === 2) {
              setFilteredData((current) => [...current, payment]);
            }
          });
          break;
        case 3:
          paymentsData.forEach((payment) => {
            if (payment.checkout_status === 3) {
              setFilteredData((current) => [...current, payment]);
            }
          });
          break;
        case 4:
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
    <Page title="Thanh toán">
      <RootStyle sx={{ px: 3, py: 2, maxWidth: '980px', bgcolor: 'red' }}>
        {/* Heading */}
        <Stack>
          <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
            Đơn mua và Số dư
          </Typography>
          <Typography variant="body2" sx={{ color: 'black', opacity: 0.8, pb: 2 }}>
            Luôn duy trì số dư trong ví để thuận tiện mua sắm
          </Typography>
        </Stack>
        {/* Main */}
        <FormProvider methods={methods} onSubmit={handleSubmit(onSubmit)}>
          <Stack direction="row" sx={{ p: 2, boxShadow: 4, borderRadius: 2 }}>
            <Stack sx={{ flex: 2 }}>
              <Stack sx={{ ml: 0.5 }}>
                <Stack direction="row">
                  <Icon icon="bi:coin" color="#eba123" fontSize="3rem" />
                  <Stack sx={{ ml: 1 }}>
                    <Typography fontSize="1.2rem" color="#eba123">
                      {userData.total_income.toLocaleString('tr-TR', {
                        style: 'currency',
                        currency: 'VND',
                      })}
                    </Typography>
                    <Typography fontStyle="italic" variant="caption" fontSize="0.9rem" sx={{ opacity: 0.7 }}>
                      Số dư hiện tại
                    </Typography>
                  </Stack>
                </Stack>
                <Typography
                  fontStyle="italic"
                  variant="body2"
                  sx={{ color: '#f44336', minWidth: '100px', opacity: 0.9, mt: 2 }}
                >
                  Mẹo: &nbsp;Số dư trong ví phải lớn hơn giá trị của sản phầm bạn muốn tham gia đấu giá
                </Typography>
              </Stack>
            </Stack>
            <Stack justifyContent="center" alignItems="flex-start" direction="row">
              <LoadingButton
                disableRipple
                color="error"
                sx={{ px: 3, textTransform: 'none' }}
                size="medium"
                type="submit"
                variant="contained"
                loading={isSubmitting}
              >
                <Icon icon="bi:coin" /> &nbsp; Nạp tiền vào ví
              </LoadingButton>
            </Stack>
          </Stack>
        </FormProvider>
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
                <Tab disableRipple label="Đã nhận" {...a11yProps(3)} />
                <Tab disableRipple label="Đã huỷ" {...a11yProps(4)} />
                <Tab disableRipple label="Trả hàng/Hoàn tiền" {...a11yProps(4)} />
              </Tabs>
            </Box>
            <TabPanel value={value} index={value}>
              {filteredData ? (
                filteredData.map((payment, index) => (
                  <Stack sx={{ boxShadow: 4, mb: 2, p: 2 }} direction="column" key={index}>
                    {/* Top side */}
                    <Stack maxHeight={20} sx={{ mb: 0.5 }} direction="row">
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
                          Xem shop
                        </Typography>
                      </Button>
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
                        {payment.checkout_status === 3 && (
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
                            onClick={() => {
                              navigate(`/auctee/user/order/?id=${payment.Id}`);
                            }}
                          >
                            Đã nhận hàng
                          </Button>
                        )}
                      </Stack>
                      {/* Total */}
                      <Stack alignItems="flex-end" flex={1}>
                        <Typography sx={{ fontSize: '0.85rem' }} variant="caption">
                          Tổng số tiền tạm tính:
                        </Typography>
                        <Typography color="#f44336">
                          {payment.before_discount.toLocaleString('tr-TR', {
                            style: 'currency',
                            currency: 'VND',
                          })}
                        </Typography>
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
                          onClick={() => {
                            navigate(`/auctee/user/order/?id=${payment.id}`);
                          }}
                        >
                          Chi tiết đơn hàng
                        </Button>
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
          <Stack>{children}</Stack>
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
