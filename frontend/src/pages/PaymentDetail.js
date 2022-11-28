import { useState, useEffect, lazy } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams, useOutletContext } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import FileBase64 from 'react-file-base64';
// material
import {
  Container,
  Avatar,
  FormGroup,
  TextField,
  RadioGroup,
  Radio,
  FormControlLabel,
  Dialog,
  DialogTitle,
  Typography,
  Stack,
  Stepper,
  Step,
  StepLabel,
  Button,
  Select,
  MenuItem,
  Divider,
} from '@mui/material';
import { LoadingButton } from '@mui/lab';
import { Icon } from '@iconify/react';
import PropTypes from 'prop-types';
import { styled, useTheme } from '@mui/material/styles';
import { Box } from '@mui/system';
import ReceiptOutlinedIcon from '@mui/icons-material/ReceiptOutlined';
import PaymentsOutlinedIcon from '@mui/icons-material/PaymentsOutlined';
import LocalShippingOutlinedIcon from '@mui/icons-material/LocalShippingOutlined';
import MoveToInboxOutlinedIcon from '@mui/icons-material/MoveToInboxOutlined';
import StarBorderPurple500OutlinedIcon from '@mui/icons-material/StarBorderPurple500Outlined';
import StepConnector, { stepConnectorClasses } from '@mui/material/StepConnector';
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

const steps = [
  {
    value: 'Xác nhận đơn hàng',
    icon: ReceiptOutlinedIcon,
  },
  {
    value: 'Chờ thanh toán',
    icon: PaymentsOutlinedIcon,
  },
  {
    value: 'Đang giao',
    icon: LocalShippingOutlinedIcon,
  },
  {
    value: 'Đã nhận',
    icon: MoveToInboxOutlinedIcon,
  },
  {
    value: 'Đánh giá',
    icon: StarBorderPurple500OutlinedIcon,
  },
];
const ColorlibConnector = styled(StepConnector)(({ theme }) => ({
  [`&.${stepConnectorClasses.alternativeLabel}`]: {
    top: 28,
  },
  [`&.${stepConnectorClasses.active}`]: {
    [`& .${stepConnectorClasses.line}`]: {
      backgroundColor: '#2dc258',
    },
  },
  [`&.${stepConnectorClasses.completed}`]: {
    [`& .${stepConnectorClasses.line}`]: {
      backgroundColor: '#2dc258',
    },
  },
  [`& .${stepConnectorClasses.line}`]: {
    height: 3,
    border: 0,
    backgroundColor: theme.palette.mode === 'dark' ? theme.palette.grey[800] : '#e0e0e0',
    borderRadius: 1,
  },
}));

const ColorlibStepIconRoot = styled('div')(({ theme, ownerState }) => ({
  backgroundColor: theme.palette.mode === 'dark' ? theme.palette.grey[700] : '#fff',
  border: '3.5px solid #e0e0e0',
  zIndex: 1,
  color: '#e0e0e0',
  width: 60,
  height: 60,
  display: 'flex',
  borderRadius: '50%',
  justifyContent: 'center',
  alignItems: 'center',
  ...(ownerState.active && {
    border: '2px solid #2dc258',
    backgroundColor: '#2dc258',
    color: 'white',
  }),
  ...(ownerState.completed && {
    border: '3.5px solid #2dc258',
    backgroundColor: 'transparent',
    color: '#2dc258',
  }),
}));

function ColorlibStepIcon(props) {
  const { active, completed, className } = props;

  const icons = {
    1: <ReceiptOutlinedIcon sx={{ fontSize: '2rem' }} />,
    2: <PaymentsOutlinedIcon sx={{ fontSize: '2rem' }} />,
    3: <LocalShippingOutlinedIcon sx={{ fontSize: '2rem' }} />,
    4: <MoveToInboxOutlinedIcon sx={{ fontSize: '2rem' }} />,
    5: <StarBorderPurple500OutlinedIcon sx={{ fontSize: '2.3rem' }} />,
  };

  return (
    <ColorlibStepIconRoot ownerState={{ completed, active }} className={className}>
      {icons[String(props.icon)]}
    </ColorlibStepIconRoot>
  );
}

ColorlibStepIcon.propTypes = {
  /**
   * Whether this step is active.
   * @default false
   */
  active: PropTypes.bool,
  className: PropTypes.string,
  /**
   * Mark the step as completed. Is passed to child components.
   * @default false
   */
  completed: PropTypes.bool,
  /**
   * The label displayed in the step icon.
   */
  icon: PropTypes.node,
};

export default function PaymentDetail() {
  const theme = useTheme();
  const navigate = useNavigate();
  const location = useLocation();

  const [searchParams, setSearchParams] = useSearchParams();
  const paymentId = searchParams.get('id');
  const productId = searchParams.get('product');
  const userData = useOutletContext();
  const [isFetching, setIsFetching] = useState(true);
  const [shopName, setShopName] = useState('');
  const [nickName, setNickName] = useState('');
  const [phoneNumber, setPhoneNumber] = useState('');
  const [avatarFile, setAvatarFile] = useState();
  const [HonorPoint, setHonorPoint] = useState(0);

  const [isMale, setIsMale] = useState(false); // 1 male : 0 female

  // All addresses
  const [userAddresses, setUserAddresses] = useState();
  const [selectedAddress, setSelectedAddress] = useState();
  const [openListAddresses, setOpenListAddresses] = useState(false);
  const [choice, setChoice] = useState();
  const handleFetchAddressData = async () => {
    await fetch('http://localhost:1001/auctee/user/addresses', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setUserAddresses(data);
          data
            .filter((address) => address.is_default === true)
            .map((item) => {
              return setSelectedAddress(item);
            });
        });
      }
      if (res.status === 401) {
        alert('You need to login first');
        navigate('/auctee/login', { replace: true });
      }
    });
  };

  // All payments
  const [paymentData, setPaymentData] = useState();
  const handleFetchPayment = async () => {
    await fetch(`http://localhost:1003/auctee/user/checkout/payment?id=${paymentId}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setPaymentData(data);
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

  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    !paymentData && handleFetchPayment();
    // eslint-disable-next-line no-unused-expressions
    !userAddresses && handleFetchAddressData();
  }, [userAddresses]);

  return !isFetching ? (
    <Page title="Chi tiết đơn hàng">
      <RootStyle sx={{ px: 3, py: 2, maxWidth: '980px' }}>
        {/* Heading */}
        <Stack>
          <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
            Chi tiết đơn hàng
          </Typography>
          <Typography variant="body2" sx={{ color: 'black', opacity: 0.8, pb: 2 }}>
            Nhập thông tin giao hàng và chọn phương thức thanh toán
          </Typography>
          <Divider />
        </Stack>
        <Stepper
          sx={{ width: '100%', my: 2 }}
          connector={<ColorlibConnector sx={{ width: '70%', ml: 1 }} />}
          activeStep={paymentData.checkout_status}
          alternativeLabel
        >
          {steps.map((item, index) => (
            <Step key={index}>
              <StepLabel sx={{ color: '#2dc258' }} StepIconComponent={ColorlibStepIcon}>
                {item.value}
              </StepLabel>
            </Step>
          ))}
        </Stepper>
        {/* Main */}
        <FormProvider methods={methods} onSubmit={handleSubmit(onSubmit)}>
          <Stack direction="row" sx={{ p: 2, boxShadow: 4, borderRadius: 2 }}>
            <Stack sx={{ flex: 2 }}>
              <Stack sx={{ ml: 0.5 }}>
                <Stack direction="row">
                  <Icon icon="material-symbols:location-on-outline-rounded" color="#f44336" fontSize="1.5rem" />
                  <Stack sx={{ ml: 1 }}>
                    <Typography fontSize="1.1rem" color="#f44336">
                      Địa chỉ nhận hàng
                    </Typography>
                    {userAddresses?.length > 0 ? (
                      <Stack alignItems="flex-end" direction="row">
                        <Typography sx={{ fontWeight: 600 }}>
                          {selectedAddress.lastname}&nbsp;
                          {selectedAddress.firstname}
                        </Typography>
                        <Typography sx={{ fontWeight: 600 }}>
                          &nbsp;&nbsp;(+84)&nbsp;{selectedAddress.phone.substring(1)}
                        </Typography>
                        <Typography variant="body1">
                          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{selectedAddress.address}, {selectedAddress.sub_district},
                          {selectedAddress.district}, {selectedAddress.province}
                        </Typography>
                        <Button
                          onClick={() => {
                            setOpenListAddresses(true);
                            setChoice(selectedAddress);
                          }}
                          disabled={paymentData.checkout_status !== 1}
                          disableRipple
                          sx={{
                            ml: 2,
                            p: 0,
                            border: 'none',
                            textTransform: 'none',
                            borderRadius: 0.4,
                            opacity: 0.85,
                            color: '#f44336',
                            '&:hover': {
                              bgcolor: 'transparent',
                              opacity: 1,
                            },
                          }}
                        >
                          Thay đổi
                        </Button>
                        {openListAddresses && (
                          <Dialog
                            open={openListAddresses}
                            sx={{ margin: 'auto', minWidth: '480px' }}
                            BackdropProps={{
                              style: { backgroundColor: 'rgba(0,0,30,0.4)' },
                              invisible: true,
                            }}
                          >
                            <DialogTitle fontWeight={500}>Chọn địa chỉ giao hàng</DialogTitle>
                            <Stack sx={{ px: 3 }}>
                              <RadioGroup
                                name="addressChoices"
                                value={choice}
                                onChange={(e) => setChoice(e.target._wrapperState.initialValue)}
                              >
                                {userAddresses?.map((address, index) => {
                                  return (
                                    <Stack maxHeight={120} direction="row" key={index}>
                                      <Stack sx={{ width: '100%', mb: 1 }}>
                                        <Stack alignItems="center" direction="row" sx={{ width: '7  0%' }}>
                                          <Typography fontSize={'1rem'} variant="body2" sx={{ color: 'black' }}>
                                            {address.lastname} {address.firstname}
                                          </Typography>
                                          <Stack sx={{ ml: 2, pl: 2, borderLeft: '1px solid grey' }}>
                                            <Typography fontSize={'0.9rem'} variant="caption" sx={{ color: 'inherit' }}>
                                              (+84) &nbsp;{address.phone.substring(1)}
                                            </Typography>
                                          </Stack>
                                        </Stack>
                                        <Typography
                                          fontSize={'0.9rem'}
                                          variant="body2"
                                          sx={{ color: 'black', opacity: 0.6 }}
                                        >
                                          {address.address}
                                        </Typography>
                                        <Typography
                                          fontSize={'0.9rem'}
                                          variant="body2"
                                          sx={{ color: 'black', opacity: 0.6 }}
                                        >
                                          {address.sub_district}, {address.district}, {address.province}
                                        </Typography>
                                        {address.is_default && (
                                          <Button
                                            fontSize="0.1rem"
                                            sx={{
                                              borderRadius: 0,
                                              color: theme.palette.background.main,
                                              textTransform: 'none',
                                              width: '20%',
                                              px: 0.5,
                                              mt: 0.5,
                                              py: 0,
                                              fontWeight: 500,
                                              border: `1px solid ${theme.palette.background.main}`,
                                            }}
                                          >
                                            Mặc định
                                          </Button>
                                        )}
                                        {userAddresses.length - index !== 1 && <Divider sx={{ mt: 2 }} />}
                                      </Stack>
                                      {/* Choices */}
                                      <Stack>
                                        <Radio
                                          sx={{ maxHeight: '100px' }}
                                          checked={choice === address}
                                          name={`address-${address.ID}`}
                                          key={index}
                                          value={address}
                                        />
                                      </Stack>
                                    </Stack>
                                  );
                                })}
                              </RadioGroup>
                              <Stack>
                                <Button
                                  onClick={() => {
                                    navigate('/auctee/user/address');
                                  }}
                                  disableRipple
                                  sx={{
                                    mx: 'auto',
                                    maxWidth: '50%',
                                    textTransform: 'none',
                                    borderRadius: 0.4,
                                    opacity: 0.7,
                                    color: 'inherit',
                                    border: '1px solid transparent',
                                    '&:hover': {
                                      border: '1px solid black',
                                      bgcolor: 'transparent',
                                      opacity: 0.9,
                                    },
                                  }}
                                >
                                  <Icon icon="material-symbols:add" /> &nbsp; Thêm địa chỉ mới
                                </Button>
                              </Stack>
                              <Stack
                                justifyContent="center"
                                alignItems="center"
                                direction="row"
                                sx={{ mt: 4, pb: 4, position: 'relative' }}
                              >
                                <Button
                                  size="medium"
                                  variant="outlined"
                                  onClick={() => setOpenListAddresses(false)}
                                  sx={{
                                    px: 1.6,
                                    position: 'absolute',
                                    right: 124,
                                    color: 'inherit',
                                    border: '1px solid white',
                                    opacity: 0.85,
                                    textTransform: 'none',
                                    '&:hover': {
                                      borderColor: 'black',
                                      opacity: 1,
                                    },
                                  }}
                                >
                                  Trở lại
                                </Button>
                                <LoadingButton
                                  onClick={() => {
                                    setSelectedAddress(choice);
                                    setOpenListAddresses(false);
                                  }}
                                  disableRipple
                                  color="error"
                                  sx={{ px: 3, position: 'absolute', right: 1, textTransform: 'none' }}
                                  size="medium"
                                  type="submit"
                                  variant="contained"
                                  loading={isSubmitting}
                                >
                                  Xác nhận
                                </LoadingButton>
                              </Stack>
                            </Stack>
                          </Dialog>
                        )}
                      </Stack>
                    ) : (
                      <Button
                        onClick={() => {
                          navigate('/auctee/user/address');
                        }}
                        disableRipple
                        sx={{
                          textTransform: 'none',
                          borderRadius: 0.4,
                          opacity: 0.7,
                          color: 'inherit',
                          '&:hover': {
                            bgcolor: 'transparent',
                            opacity: 0.9,
                          },
                        }}
                      >
                        <Icon icon="material-symbols:add" /> &nbsp; Thêm địa chỉ mới
                      </Button>
                    )}
                  </Stack>
                </Stack>
                <Typography
                  fontStyle="italic"
                  variant="body2"
                  sx={{ color: '#f44336', minWidth: '100px', opacity: 0.9, mt: 2 }}
                >
                  Mẹo: &nbsp;Địa chỉ càng chính xác sẽ giúp bạn nhận hàng càng nhanh
                </Typography>
              </Stack>
            </Stack>
            {/* <Stack justifyContent="center" alignItems="flex-start" direction="row">
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
            </Stack> */}
          </Stack>
        </FormProvider>
      </RootStyle>
    </Page>
  ) : (
    <>Có lỗi xảy ra</>
  );
}
