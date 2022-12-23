import { useState, useEffect, lazy } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams, useOutletContext } from 'react-router-dom';
// material
import {
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
  Input,
  Divider,
  useMediaQuery,
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
import ProductionQuantityLimitsOutlinedIcon from '@mui/icons-material/ProductionQuantityLimitsOutlined';
import { Dataset } from '@mui/icons-material';

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

const stepsCheckedOut = [
  {
    value: 'Xác nhận đơn hàng',
    icon: ReceiptOutlinedIcon,
  },
  {
    value: 'Đã thanh toán',
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
const stepsForCancel = [
  {
    value: 'Đơn hàng đã đặt',
    icon: ReceiptOutlinedIcon,
  },
  {
    value: 'Đơn hàng đã bị huỷ',
    icon: ProductionQuantityLimitsOutlinedIcon,
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
function ColorlibStepIconCancel(props) {
  const { active, completed, className } = props;

  const icons = {
    1: <ReceiptOutlinedIcon sx={{ fontSize: '2rem' }} />,
    2: <ProductionQuantityLimitsOutlinedIcon sx={{ fontSize: '2rem' }} />,
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

export default function OrderDetail() {
  const theme = useTheme();
  const navigate = useNavigate();

  const [searchParams, setSearchParams] = useSearchParams();
  const paymentId = searchParams.get('id');
  const resultCode = searchParams.get('resultCode');

  const userData = useOutletContext();
  const [isFetching, setIsFetching] = useState(true);

  const [noteData, setNoteData] = useState('');

  // All addresses
  const [userAddresses, setUserAddresses] = useState();
  const [selectedAddress, setSelectedAddress] = useState();
  const [openListAddresses, setOpenListAddresses] = useState(false);
  const [paymentMethod, setPaymentMethod] = useState('COD');
  const [paymentStatus, setPaymentStatus] = useState('');

  const [stepStatus, setStepStatus] = useState(0);

  const handleMoMoCheckOutStatus = async () => {
    console.log(resultCode);
    if (resultCode === '0') {
      const payload = {
        total: parseFloat(paymentData?.before_discount + shippingFee),
        shipping_value: parseFloat(shippingFee),
        id: paymentData?.id,
      };
      await fetch(`http://localhost:8080/auctee/user/update/momo-payment`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        mode: 'cors',
        body: JSON.stringify(payload),
      }).then((res) => {
        if (res.status === 200) {
          setPaymentStatus('Đang vận chuyển');
          navigate(0);
        }
        if (res.status === 401) {
          alert('You need to login first');
          navigate('/auctee/login', { replace: true });
        }
      });
    }
    if (resultCode === '1006' && paymentData.payment_method === 'MOMO') {
      alert('Giao dịch MoMo không thành công');
    }
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
  const [choice, setChoice] = useState();

  const fullScreen = useMediaQuery(theme.breakpoints.down('md'));
  const [openConfirmCancel, setOpenConfirmCancel] = useState(false);
  const handleCancel = async (id) => {
    await fetch(`http://localhost:8080/auctee/user/checkout/cancel-payment?id=${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',

      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        setOpenListAddresses(false);
        navigate(0);
      }
      if (res.status === 400) {
        setOpenListAddresses(true);
      }
      if (res.status === 401) {
        alert('You need to login first');
        navigate('/auctee/login', { replace: true });
      }
    });
  };

  const [currAddress, setCurrAddress] = useState();
  // All payments
  const [paymentData, setPaymentData] = useState();
  const handleFetchPayment = async () => {
    await fetch(`http://localhost:8080/auctee/user/sale/payment?id=${paymentId}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setPaymentData(data);
          setCurrAddress(data.address_id);
          setSelectedAddress({
            province: data.province,
            sub_district: data.sub_district,
            district: data.district,
            address: data.address,
            type_address: data.type_address,
            firstname: data.firstname,
            lastname: data.lastname,
            phone: data.phone,
          });
          // setTotalBill(data.before_discount + data.shipping_value);
          setIsFetching(false);
          setPaymentMethod(data.payment_method || 'COD');
        });
      }
      if (res.status === 401) {
        alert('You need to login first');
        setIsFetching(true);
        navigate('/auctee/login', { replace: true });
      }
    });
  };
  // const [totalBill, setTotalBill] = useState();
  const [shippingFee, setShippingFee] = useState();
  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    selectedAddress?.province === 'Hồ Chí Minh' || currAddress === 'Hồ Chí Minh'
      ? setShippingFee(25000)
      : setShippingFee(35000);
  }, [selectedAddress, currAddress]);

  const handleStatus = () => {
    switch (paymentData.checkout_status) {
      case 1:
        setPaymentStatus('Chưa thanh toán');
        break;
      case 2:
        if (paymentData.total > 0) {
          setPaymentStatus('Đã thanh toán');
        }
        break;
      case 3:
        if (paymentData.shipping_status === 1 || paymentData.shipping_status === 2) {
          setPaymentStatus('Đang vận chuyển');
          break;
        } else {
          setPaymentStatus('Đã nhận hàng');
          break;
        }
      case 4:
        setPaymentStatus('Đã huỷ');
        break;
      case 5:
        setPaymentStatus('Hoàn thành');
        break;
      default:
        setPaymentStatus('Chưa thanh toán');
    }
  };

  const handleStep = () => {
    switch (paymentData.checkout_status) {
      case 1:
        if (paymentData.total === 0) {
          setStepStatus(1);
        }
        break;
      case 2: // đang giao
        setStepStatus(1);
        break;
      case 3: // đã nhận
        if (paymentData.shipping_status === 3 && paymentData.total !== 0) {
          setStepStatus(3);
        } else {
          setStepStatus(2);
        }
        break;
      case 4: // huỷ
        setStepStatus(1);
        break;
      case 5:
        setStepStatus(5);
        break;
      default:
        setStepStatus(2);
    }
  };
  const [shippingStatus, setShippingStatus] = useState();

  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    !paymentData && handleFetchPayment();
    // eslint-disable-next-line no-unused-expressions
    if (paymentData) {
      handleStep();
      handleStatus();
      if ((resultCode === '0' || resultCode === '1006') && paymentData.checkout_status < 3) {
        handleMoMoCheckOutStatus();
      }
      if (paymentData.shipping_status === 1) {
        setShippingStatus('Chờ shipper lấy hàng');
      }
      if (paymentData.shipping_status === 2 && paymentData.checkout_status !== 4) {
        setShippingStatus('Đã giao cho đơn vị vận chuyển');
      }
      if (paymentData.shipping_value !== 0) {
        setSelectedAddress({
          province: paymentData.province,
          sub_district: paymentData.sub_district,
          district: paymentData.district,
          address: paymentData.address,
          type_address: paymentData.type_address,
          firstname: paymentData.firstname,
          lastname: paymentData.lastname,
          phone: paymentData.phone,
        });
      }
    }
  }, [paymentData]);

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
        {/* Stepper */}
        {paymentData.checkout_status === 4 ? (
          <Stepper
            sx={{ width: '100%', my: 2 }}
            connector={<ColorlibConnector sx={{ width: '100%', ml: 1 }} />}
            activeStep={1}
            alternativeLabel
          >
            {stepsForCancel.map((item, index) => (
              <Step key={index}>
                <StepLabel sx={{ color: '#2dc258' }} StepIconComponent={ColorlibStepIconCancel}>
                  {item.value}
                </StepLabel>
              </Step>
            ))}
          </Stepper>
        ) : (
          <Stepper
            sx={{ width: '100%', my: 2 }}
            connector={<ColorlibConnector sx={{ width: '70%', ml: 1 }} />}
            activeStep={stepStatus}
            alternativeLabel
          >
            {paymentData.total === 0
              ? steps.map((item, index) => (
                  <Step key={index}>
                    <StepLabel sx={{ color: '#2dc258' }} StepIconComponent={ColorlibStepIcon}>
                      {item.value}
                    </StepLabel>
                  </Step>
                ))
              : stepsCheckedOut.map((item, index) => (
                  <Step key={index}>
                    <StepLabel sx={{ color: '#2dc258' }} StepIconComponent={ColorlibStepIcon}>
                      {item.value}
                    </StepLabel>
                  </Step>
                ))}
          </Stepper>
        )}
        {shippingStatus && (
          <Stack sx={{ mt: -2, pb: 2, mx: 'auto' }}>
            <Typography fontStyle="italic" variant="caption" color="primary">
              {shippingStatus}
            </Typography>
          </Stack>
        )}
        {/* Main */}
        <Stack direction="row" sx={{ p: 2, boxShadow: 4 }}>
          {/* Address information */}
          <Stack sx={{ flex: 2 }}>
            <Stack sx={{ ml: 0.5 }}>
              <Stack direction="row">
                <Icon icon="material-symbols:location-on-outline-rounded" color="#F62217" fontSize="1.4rem" />
                <Stack sx={{ ml: 1 }}>
                  <Typography fontSize="1rem" color="#F62217">
                    Địa chỉ nhận hàng
                  </Typography>
                  {selectedAddress ? (
                    <Stack alignItems="flex-end" direction="row">
                      <Typography sx={{ fontWeight: 600, fontSize: '0.95rem' }}>
                        {selectedAddress?.lastname}&nbsp;
                        {selectedAddress?.firstname}
                      </Typography>
                      <Typography sx={{ fontWeight: 600, fontSize: '0.95rem' }}>
                        &nbsp;&nbsp;(+84)&nbsp;{selectedAddress?.phone.slice(1)}
                      </Typography>
                      <Typography variant="body1" sx={{ fontSize: '0.95rem' }}>
                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{selectedAddress?.address}, {selectedAddress?.sub_district},
                        {selectedAddress?.district}, {selectedAddress?.province}
                      </Typography>
                    </Stack>
                  ) : (
                    <Stack>Chưa có thông tin giao hàng</Stack>
                  )}
                </Stack>
              </Stack>
            </Stack>
          </Stack>
        </Stack>
        {/* Product information */}
        <Stack sx={{ boxShadow: 4, my: 2, p: 2 }} direction="column">
          {/* Body */}
          <Stack justifyContent="space-between" direction="row" sx={{ display: 'flex', height: '100%' }}>
            <Stack flex={4} direction="column" justifyContent="space-between">
              {/* Image */}
              <Stack flex={2} direction="row">
                <Stack>
                  <ProductImgStyle alt={paymentData.product_id} src={paymentData.image_path} />
                </Stack>
                {/* Name */}
                <Stack sx={{ mx: 2 }}>
                  <Typography variant="caption" sx={{ textTransform: 'uppercase', fontSize: '1rem' }}>
                    {paymentData.product_name}
                  </Typography>
                  <Typography sx={{ mt: 1 }}>x{paymentData.quantity}</Typography>
                </Stack>
              </Stack>
              {/* Note */}
              <Stack sx={{ ml: 1.5 }} alignItems="center" direction="row">
                <Typography
                  color="primary"
                  variant="caption"
                  sx={{ mr: 3, fontSize: '0.9rem', whiteSpace: 'nowrap', fontStyle: 'italic' }}
                >
                  Lời nhắn:&nbsp;
                </Typography>
                <Input
                  disabled={paymentData.checkout_status !== 1}
                  value={noteData}
                  onChange={(e) => setNoteData(e.target.value)}
                  fullWidth
                  disableUnderline
                  placeholder={`${paymentData.note ? paymentData.note : 'Lưu ý cho người bán'}`}
                  inputProps={{
                    sx: {
                      width: '70%',
                      px: 2,
                      border: '1px solid grey',
                      '&::placeholder': {
                        opacity: 0.7,
                        color: 'black',
                        fontWeight: 200,
                        fontSize: '0.8rem',
                      },
                    },
                  }}
                />
              </Stack>
            </Stack>
            {/* Total */}
            <Stack justifyContent="space-between" sx={{ height: '100%' }} alignItems="flex-end" flex={2}>
              <Stack width="105%" justifyContent="flex-end" direction="row" alignItems="flex-end">
                <Typography sx={{ fontSize: '0.85rem' }} variant="caption">
                  ID: &nbsp;
                </Typography>
                <Typography sx={{ fontSize: '0.85rem' }} variant="caption">
                  {paymentData.id}
                </Typography>
                <Typography sx={{ fontSize: '0.85rem', borderLeft: '1px solid grey', pl: 1, ml: 1 }} variant="caption">
                  Trạng thái: &nbsp;&nbsp;&nbsp;
                </Typography>
                <Typography
                  color={`${paymentData.checkout_status < 2 ? '#f5b70c' : '#2dc258'}`}
                  sx={{ fontSize: '0.85rem' }}
                  variant="caption"
                >
                  {paymentStatus}
                </Typography>
              </Stack>
              <Stack width="65%" justifyContent="space-between" direction="row" alignItems="center">
                <Typography variant="caption" sx={{ fontSize: '0.9rem', opacity: 0.8 }}>
                  Vận chuyển:
                </Typography>
                <Typography variant="caption" sx={{ fontSize: '1rem', opacity: 0.7, ml: 2 }}>
                  {shippingFee.toLocaleString('tr-TR', {
                    style: 'currency',
                    currency: 'VND',
                  })}
                </Typography>
              </Stack>
              <Stack width="65%" justifyContent="space-between" direction="row" alignItems="center">
                <Typography variant="caption" sx={{ fontSize: '0.9rem', opacity: 0.8 }}>
                  Tồng cộng:
                </Typography>
                <Typography sx={{ ml: 2 }} color="#F62217">
                  {(paymentData.before_discount + shippingFee).toLocaleString('tr-TR', {
                    style: 'currency',
                    currency: 'VND',
                  })}
                </Typography>
              </Stack>
              {/* Payment Methods */}
              <Stack direction="column" sx={{ mb: 0.5 }} width="77%">
                <Typography variant="caption" sx={{ fontSize: '0.9rem' }}>
                  Phương thức thanh toán:
                </Typography>
                <RadioGroup
                  row
                  aria-labelledby="demo-radio-buttons-group-label"
                  value={paymentMethod}
                  onChange={(e) => setPaymentMethod(e.target.value)}
                  name="radio-buttons-group"
                >
                  <FormControlLabel
                    componentsProps={{
                      typography: {
                        variant: 'caption',
                        color: 'inherit',
                        border: '1px solid grey',
                        px: 1,
                        py: 0.15,
                        borderRadius: 0.5,
                      },
                    }}
                    value="COD"
                    control={
                      <Radio
                        sx={{
                          p: 0.7,
                          color: '#F62217',
                          '&.Mui-checked': {
                            color: '#F62217',
                          },
                        }}
                        disabled={paymentData.checkout_status !== 1}
                        size="small"
                      />
                    }
                    label="COD"
                  />
                  <FormControlLabel
                    componentsProps={{
                      typography: {
                        variant: 'caption',
                        color: 'white',
                        bgcolor: '#a50064',
                        px: 1,
                        py: 0.2,
                        borderRadius: 0.5,
                      },
                    }}
                    value="MOMO"
                    control={
                      <Radio
                        sx={{
                          p: 0.7,
                          color: '#F62217',
                          '&.Mui-checked': {
                            color: '#F62217',
                          },
                        }}
                        disabled={paymentData.checkout_status !== 1}
                        size="small"
                      />
                    }
                    label="Ví MoMo E-Wallet"
                  />
                </RadioGroup>
              </Stack>
              <Stack justifyContent="flex-end" direction="row">
                <Button
                  disabled={paymentData.checkout_status !== 1}
                  size="medium"
                  variant="outlined"
                  disableRipple
                  sx={{
                    whiteSpace: 'nowrap',
                    border: '1px solid black',

                    borderRadius: 0.4,
                    color: 'inherit',
                    px: 1.5,
                    textTransform: 'none',
                    '&:hover': {
                      bgcolor: 'transparent',
                      border: '1px solid black',
                      opacity: 0.8,
                    },
                  }}
                  onClick={() => {
                    setOpenConfirmCancel(true);
                  }}
                >
                  Huỷ đơn
                </Button>
                {/* Dialog Delete */}
                <Dialog
                  sx={{ margin: 'auto', minWidth: '480px' }}
                  BackdropProps={{
                    style: { backgroundColor: 'rgba(0,0,30,0.2)' },
                    invisible: true,
                  }}
                  fullScreen={fullScreen}
                  open={openConfirmCancel}
                >
                  <Stack sx={{ p: 3, overflow: 'hidden' }}>Bạn có chắc muốn huỷ đơn hàng này?</Stack>
                  <Stack sx={{ px: 3 }}>
                    <Typography color="error" sx={{ fontSize: '0.9rem' }} variant="caption" fontStyle="italic">
                      (Lưu ý: Nếu bạn huỷ ngay bây giờ sẽ bị trừ 5 điềm uy tín)
                    </Typography>
                  </Stack>
                  <Stack sx={{ p: 2 }} justifyContent="flex-end" direction="row" alignItems="center">
                    <Button
                      disableRipple
                      sx={{
                        color: 'inherit',
                        bgcolor: 'transparent',
                        opacity: 0.85,
                        border: '1px solid white',
                        textTransform: 'none',
                        '&:hover': {
                          bgcolor: 'transparent',
                          opacity: 1,
                          border: '1px solid black',
                        },
                      }}
                      onClick={() => setOpenConfirmCancel(false)}
                    >
                      Trở lại
                    </Button>
                    <Button
                      disableRipple
                      color="error"
                      variant="contained"
                      sx={{
                        ml: 1,
                        color: 'white',
                        bgcolor: '#F62217',
                        textTransform: 'none',
                      }}
                      onClick={() => handleCancel(paymentData.id)}
                      autoFocus
                    >
                      Xác nhận huỷ
                    </Button>
                  </Stack>
                </Dialog>
                {paymentData.shipping_status === 1 && (
                  <Stack sx={{}} width="100%" direction="row" justifyContent="flex-end">
                    <Button
                      color="error"
                      size="medium"
                      variant="contained"
                      disableRipple
                      sx={{
                        borderRadius: 0.4,
                        bgcolor: '#F62217',
                        color: 'white',
                        px: 1.5,
                        textTransform: 'none',
                      }}
                      onClick={() => handleConfirmDelivery(paymentData)}
                    >
                      Giao cho đơn vị vận chuyển
                    </Button>
                  </Stack>
                )}
                {paymentData.shipping_status === 2 && (
                  <Stack width="100%" direction="row" justifyContent="flex-end">
                    <Button
                      disabled={paymentData.checkout_status === 4}
                      color="error"
                      size="medium"
                      variant="contained"
                      disableRipple
                      sx={{
                        borderRadius: 0.4,
                        bgcolor: '#F62217',
                        color: 'white',
                        px: 1.5,
                        textTransform: 'none',
                      }}
                      onClick={() => handleShopCancel(paymentData)}
                    >
                      Không giao được hàng
                    </Button>
                  </Stack>
                )}
                {((paymentData.checkout_status === 3 && paymentData.shipping_status === 3) ||
                  paymentData.checkout_status === 5) && (
                  <Stack sx={{}} width="100%" direction="row" justifyContent="flex-end">
                    <Button
                      disabled={paymentData.checkout_status === 5}
                      color="error"
                      size="medium"
                      variant="contained"
                      disableRipple
                      sx={{
                        borderRadius: 0.4,
                        bgcolor: '#F62217',
                        color: 'white',
                        px: 1.5,
                        textTransform: 'none',
                        '&:disabled': {
                          backgroundColor: `${paymentData.checkout_status === 5 && '  '}`,
                          color: `${paymentData.checkout_status === 5 && 'white'}`,
                          opacity: `${paymentData.checkout_status === 5 && '0.8'}`,
                        },
                      }}
                      onClick={() => handleConfirmDone(paymentData)}
                    >
                      {paymentData.checkout_status === 5 ? 'Đã hoàn thành' : 'Xác nhận hoàn thành đơn hàng'}
                    </Button>
                  </Stack>
                )}
              </Stack>
            </Stack>
          </Stack>
        </Stack>
      </RootStyle>
    </Page>
  ) : (
    <>Không tìm thấy đơn hàng</>
  );
}

const ProductImgStyle = styled('img')({
  width: '100px',
  height: '105px',
  objectFit: 'cover',
});
