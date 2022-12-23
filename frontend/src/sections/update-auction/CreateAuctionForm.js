import { useState, useEffect, lazy, Suspense, useContext } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import Select from 'react-select';

import dayjs from 'dayjs';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
// material
import {
  Dialog,
  Checkbox,
  DialogTitle,
  Typography,
  Stack,
  Button,
  Divider,
  Avatar,
  TextField,
  FormControlLabel,
  Switch,
} from '@mui/material';
import AddPhotoAlternateOutlinedIcon from '@mui/icons-material/AddPhotoAlternateOutlined';
import { LoadingButton } from '@mui/lab';
import { styled, useTheme } from '@mui/material/styles';

import { Icon } from '@iconify/react';
import { FormProvider, RHFTextField } from '../../components/hook-form';
import { ReloadContext } from '../../utils/Context';

const customStyles = {
  option: (provided, state) => ({
    ...provided,
    color: state.isSelected && '#F62217',
    backgroundColor: state.isSelected && 'white',
  }),
  control: (base, state) => ({
    ...base,
    minWidth: '114px !important',
    maxWidth: '114px !important',
    whiteSpace: 'nowrap',
    borderRadius: 8,
    marginBottom: '4%',
    boxShadow: 'none',
    border: state.isFocused && '1px solid #F62217 !important',
    '&:hover': {
      border: '1px solid black',
    },
  }),
  menu: (base) => ({
    ...base,
    marginTop: '-4%',
    zIndex: 100,
  }),
  menuList: (base) => ({
    ...base,
    marginTop: 0,
    maxHeight: '200px',
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
  }),
};

export default function CreateProductForm() {
  const theme = useTheme();
  const navigate = useNavigate();
  const { isReloading, setIsReloading } = useContext(ReloadContext);

  const [open, setOpen] = useState(false);
  const [openCreateOption, setOpenOption] = useState(false);
  const [openUpdateOption, setOpenUpdateOption] = useState(false);

  // product information
  const [proId, setProId] = useState('');
  const [proName, setProName] = useState('');
  const [minPrice, setMinPrice] = useState('');
  const [proDescription, setDescription] = useState('');
  const [auctionStep, setAuctionStep] = useState('');
  const [expectPrice, setExpectPrice] = useState('');
  const [startDate, setStartDate] = useState();
  const [endDate, setEndDate] = useState();
  const [imageAuction, setImageAuction] = useState();
  const [isActive, setIsActive] = useState(false);

  const [images, setImages] = useState([]);

  const [isFetching, setIsFetching] = useState(true);
  const [errMessage, setErrorMessage] = useState('');

  const defaultValues = {
    id: '',
    name: '',
    min_price: '',
    expect_price: '',
    quantity: '',
    description: '',
  };

  const methods = useForm({
    defaultValues,
  });
  const {
    handleSubmit,
    formState: { isSubmitting },
  } = methods;

  const [error, setError] = useState(false);
  const onSubmit = async () => {
    console.log(startDate.$d.toISOString().slice(0, 19).replace('T', ' '));
    const payload = {
      product_id: proId,
      quantity: 1,
      start_time: startDate.$d.toISOString().slice(0, 19).replace('T', ' '),
      end_time: endDate.$d.toISOString().slice(0, 19).replace('T', ' '),
      is_active: isActive,
      price_per_step: parseFloat(auctionStep),
      image_path: imageAuction,
    };

    await fetch('http://localhost:8080/auctee/user/auction', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
      credentials: 'include',

      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        setError(false);
      }
      setIsReloading(true);
      setOpen(false);
      if (res.status === 400 || res.status === 409) {
        setError(true);
        setErrorMessage('Thời gian không hợp lệ');
        setOpen(true);
      }
    });
  };

  const [productsData, setProductData] = useState();
  const fetchProductData = async () => {
    await fetch('http://localhost:8080/auctee/products', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',

      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setProductData(data);
          setIsFetching(false);
        });
      }
      if (res.status === 401) {
        setIsFetching(true);
        alert('You need to login first');
        navigate('/auctee/login', { replace: true });
      }
    });
  };
  const [idList, setIdList] = useState([]);

  const handleClickOpen = () => {
    if (productsData?.length > 0) {
      setStartDate(dayjs());
      setEndDate(dayjs());

      /* eslint-disable-next-line no-plusplus */
      for (let i = 0; i < productsData?.length; i++) {
        idList.push({ value: productsData[i].id, label: productsData[i].id });
      }
      setProId(productsData[0]?.id);
      setProName(productsData[0]?.name);
      setMinPrice(productsData[0]?.min_price);
      setExpectPrice(productsData[0]?.expect_price);
      setDescription(productsData[0]?.description);
      setImages(productsData[0]?.product_images);
      setImageAuction(productsData[0]?.product_images[0].path);
      setOpen(true);
    } else {
      alert('Cần thêm sản phẩm trước khi mở đấu giá');
    }
  };

  const handleChangeId = async (id) => {
    await fetch(`http://localhost:8080/auctee/product/detail?id=${id}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',

      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        res.json().then((data) => {
          setProId(data.id);
          setProName(data.name);
          setMinPrice(data.min_price);
          setExpectPrice(data.expect_price);
          setDescription(data.description);
          setImages(data.product_images);
          setImageAuction(data.product_images[0].path);
        });
      }
    });
  };

  const handleClickImage = (path) => {
    setImageAuction(path);
  };

  useEffect(() => {
    setIsReloading(false);
    // eslint-disable-next-line no-unused-expressions
    !productsData && fetchProductData();
  }, [isFetching, isReloading]);

  return (
    <>
      <Stack justifyContent="space-between" alignItems="center" direction="row" sx={{ maxHeight: '100%', pb: 1 }}>
        <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
          Tất cả phiên đấu giá
        </Typography>
        <Button
          onClick={handleClickOpen}
          sx={{
            my: 1,
            px: '20px !important',
            color: 'white',
            bgcolor: theme.palette.background.main,
            fontWeight: 500,
            textTransform: 'none',
          }}
          color="error"
          variant="contained"
          component="label"
        >
          <Icon icon="akar-icons:plus" />
          &nbsp; Thêm phiên đấu giá mới
        </Button>
      </Stack>
      {/* Form create auction */}
      <Dialog
        open={open}
        fullWidth
        maxWidth="md"
        BackdropProps={{
          style: { backgroundColor: 'rgba(0,0,30,0.4)' },
          invisible: true,
        }}
      >
        <DialogTitle fontWeight={500}>Thông tin cơ bản</DialogTitle>
        <FormProvider methods={methods} onSubmit={handleSubmit(onSubmit)}>
          <Stack direction="row">
            {/* Left */}
            <Stack sx={{ px: 3, width: '60%', pr: 0 }}>
              {/* Basic infor */}
              <Stack justifyContent="space-between" direction="row" sx={{ pb: 2 }}>
                <Select
                  styles={customStyles}
                  name="prodId"
                  options={idList}
                  defaultValue={proId}
                  onChange={(item) => handleChangeId(item.label)}
                  /*  eslint no-unneeded-ternary: "error" */
                  placeholder={`${proId ? proId : 'Mã sp'}`}
                />
                <RHFTextField
                  color="error"
                  required
                  label="Tên sản phẩm"
                  name="name"
                  type="text"
                  value={proName || ''}
                  size="small"
                  variant="outlined"
                  sx={{ width: '73%' }}
                  inputProps={{ readOnly: true }}
                />
              </Stack>
              {/* Step */}
              <Stack justifyContent="space-between" alignItems="center" direction="row" sx={{ pb: 2 }}>
                <RHFTextField
                  InputProps={{
                    inputProps: {
                      min: 1000,
                    },
                  }}
                  color="error"
                  required
                  label="Bước giá"
                  name="step"
                  type="number"
                  value={auctionStep || ''}
                  onChange={(e) => setAuctionStep(e.target.value)}
                  size="small"
                  variant="outlined"
                  sx={{ width: '22%' }}
                />
                <Divider orientation="vertical" flexItem sx={{ mx: 0.8 }} />
                <RHFTextField
                  color="error"
                  required
                  label="Giá khởi điểm"
                  name="minprice"
                  type="number"
                  value={minPrice || ''}
                  size="small"
                  variant="outlined"
                  sx={{ width: '27%' }}
                  inputProps={{ readOnly: true }}
                />
                <Typography variant="body1">VNĐ</Typography>
                -
                <RHFTextField
                  color="error"
                  required
                  label="Giá tối đa"
                  name="expectprice"
                  type="text"
                  value={expectPrice || ''}
                  size="small"
                  variant="outlined"
                  sx={{ width: '27%' }}
                  inputProps={{ readOnly: true }}
                />
                <Typography variant="body1">VNĐ</Typography>
              </Stack>
              {/* Start-End date */}
              <Stack justifyContent="space-between" alignItems="center" direction="row" sx={{ pb: 2 }}>
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                  <DateTimePicker
                    renderInput={(props) => <TextField {...props} sx={{ display: 'flex', flex: 1.5 }} />}
                    label="Ngày bắt đầu"
                    value={startDate || dayjs()}
                    onChange={(newValue) => {
                      setStartDate(newValue);
                    }}
                  />
                  <DateTimePicker
                    renderInput={(props) => <TextField {...props} sx={{ display: 'flex', flex: 1.5, ml: 1.5 }} />}
                    label="Ngày kết thúc"
                    value={endDate || dayjs()}
                    onChange={(newValue) => {
                      setEndDate(newValue);
                    }}
                  />
                </LocalizationProvider>
                {/* Active  */}
                <FormControlLabel
                  sx={{ display: 'flex', flex: 1.2, fontSize: '0.2rem', ml: 1.5, mr: 0 }}
                  control={<Switch color="success" onChange={(e) => setIsActive(e.target.checked)} />}
                  label="Kích hoạt"
                />
              </Stack>
              {/* Description */}
              <Stack direction="row" sx={{ flex: 2, pb: 2 }}>
                <RHFTextField
                  color="error"
                  required
                  label="Mô tả chi tiết"
                  name="description"
                  type="text"
                  value={proDescription || ''}
                  onChange={(e) => setDescription(e.target.value)}
                  multiline
                  rows={4}
                  variant="outlined"
                  inputProps={{ readOnly: true }}
                />
              </Stack>
              <Stack alignItems="center" direction="row" sx={{ height: 3 }}>
                {error && (
                  <Typography variant="body2" color="error">
                    {errMessage}
                  </Typography>
                )}
              </Stack>
            </Stack>
            {/* Right */}
            <Stack direction="row" sx={{ px: 3, width: '37.2%' }}>
              <Divider orientation="vertical" />
              <Stack
                justifyContent="flex-start"
                alignItems="flex-start"
                direction="row"
                sx={{ ml: 3, pr: 0, minWidth: '100%', mb: 1 }}
              >
                {/* Images */}
                <Stack alignItems="center">
                  <Avatar
                    variant="square"
                    key={imageAuction}
                    alt="main_image"
                    src={imageAuction}
                    sx={{ width: 234, height: 257, mr: `${images?.length > 0 ? '0' : '56px'}` }}
                  >
                    <AddPhotoAlternateOutlinedIcon sx={{ fontSize: '4rem', fontWeight: 400 }} />
                  </Avatar>
                  <Typography variant="subtitle1" sx={{ mt: 2 }}>
                    Chọn ảnh bìa cho sản phẩm
                  </Typography>
                </Stack>
                <Stack sx={{ flexWrap: 'wrap' }} direction="column">
                  {images.map((item, index) => (
                    <Stack key={index} sx={{ mx: 0.4, mt: '-2px' }}>
                      <Avatar
                        onClick={() => handleClickImage(item.path)}
                        variant="square"
                        alt="image"
                        src={item.path}
                        sx={{
                          width: 50,
                          height: 50,
                          borderRadius: 0.5,
                          mb: 0.6,
                          border:
                            imageAuction === item.path
                              ? `2px solid ${theme.palette.background.main}`
                              : '2px solid transparent',
                          '&:hover': {
                            border: `2px solid ${theme.palette.background.main}`,
                          },
                        }}
                      />
                    </Stack>
                  ))}
                </Stack>
              </Stack>
            </Stack>
          </Stack>
          {/* Buttons */}
          <Stack
            justifyContent="center"
            alignItems="center"
            direction="row"
            sx={{ mt: 2, pb: 4, position: 'relative' }}
          >
            <Button
              size="medium"
              variant="outlined"
              onClick={() => setOpen(false)}
              sx={{
                px: 1.6,
                position: 'absolute',
                right: 116,
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
              disableRipple
              color="error"
              sx={{ px: 3, position: 'absolute', right: 24, bgcolor: '#F62217' }}
              size="medium"
              type="submit"
              variant="contained"
              loading={isSubmitting}
            >
              Lưu
            </LoadingButton>
          </Stack>
        </FormProvider>
      </Dialog>
    </>
  );
}
