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
  RadioGroup,
  Radio,
  Typography,
  Stack,
  Button,
  Divider,
  Box,
  alertTitleClasses,
  FormControlLabel,
  IconButton,
  Avatar,
  TextField,
  Switch,
  Tooltip,
} from '@mui/material';

import useMediaQuery from '@mui/material/useMediaQuery';
import { LoadingButton } from '@mui/lab';
import AddPhotoAlternateOutlinedIcon from '@mui/icons-material/AddPhotoAlternateOutlined';
import HighlightOffRoundedIcon from '@mui/icons-material/HighlightOffRounded';
import { styled, useTheme } from '@mui/material/styles';
import { Icon } from '@iconify/react';
import { FormProvider, RHFTextField } from '../../components/hook-form';
import useLocationForm from '../update-address/useLocationForm';
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

export default function UpdateAddressForm({ auction, handleDelete, index }) {
  const theme = useTheme();
  const navigate = useNavigate();

  const { isReloading, setIsReloading } = useContext(ReloadContext);

  const [open, setOpen] = useState(false);
  const [openCreateOption, setOpenOption] = useState(false);
  const [openUpdateOption, setOpenUpdateOption] = useState(false);

  // auction information
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

  const stringToBoolean = (value) => {
    return String(value) === '1' || String(value).toLowerCase() === 'true';
  };
  const defaultValues = {};

  const methods = useForm({
    defaultValues,
  });
  const {
    handleSubmit,
    formState: { isSubmitting },
  } = methods;

  // update auction
  const onSubmit = async () => {
    const payload = {
      image_path: imageAuction,
      is_active: stringToBoolean(isActive),
      price_per_step: parseFloat(auctionStep),
      start_time:
        typeof startDate === 'object'
          ? startDate.$d.toISOString().slice(0, 19).replace('T', ' ')
          : startDate.slice(0, 19).replace('T', ' '),
      end_time:
        typeof endDate === 'object'
          ? endDate.$d.toISOString().slice(0, 19).replace('T', ' ')
          : endDate.slice(0, 19).replace('T', ' '),
    };

    await fetch(`http://localhost:8080/auctee/user/auction/detail?id=${auction?.Id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
      credentials: 'include',
      mode: 'cors',
    }).then((res) => {
      if (res.status === 200) {
        setError(false);
        setOpen(false);
        setIsReloading(true);
      }
      if (res.status === 400) {
        setError(true);
        setErrorMessage('Thời gian không hợp lệ');
        setOpen(true);
        setIsReloading(false);
      }
      if (res.status === 401) {
        alert('You need to login first');
        navigate('/auctee/login', { replace: true });
      }
    });
  };
  const [error, setError] = useState(false);

  const fullScreen = useMediaQuery(theme.breakpoints.down('md'));

  const [openDelForm, setOpenDelForm] = useState(false);
  const handleClickOpen = () => {
    fetchProduct(auction?.product_id);
    setStartDate(auction?.start_time);
    setEndDate(auction?.end_time);
    setProName(auction?.product_name);
    setAuctionStep(auction?.price_per_step);
    setImageAuction(auction?.image_path);
    setIsActive(auction.is_active);
    setOpen(true);
  };

  const fetchProduct = async (id) => {
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
        });
      }
    });
  };

  const handleClickImage = (path) => {
    setImageAuction(path);
  };
  // delete form
  const handleClickOpenDelForm = () => {
    setOpenDelForm(true);
  };

  const handleCloseDelForm = () => {
    setOpenDelForm(false);
  };

  return (
    <>
      <Stack key={index} direction="row">
        <Tooltip enterDelay={700} title="Lưu ý: không thể cập nhật phiên đấu giá đã có người tham gia">
          <span>
            <Button
              onClick={handleClickOpen}
              sx={{
                borderRadius: 0,
                maxWidth: '80px',
                p: 0,
                bgcolor: 'transparent',
                border: 'none',
                fontSize: '0.9rem',
                color: 'green',
                fontWeight: 400,
                textTransform: 'none',
                '&:hover': {
                  bgcolor: 'transparent',
                  textDecoration: 'underline',
                },
              }}
            >
              Cập nhật
            </Button>
          </span>
        </Tooltip>
        {/* Dialog Update */}
        <Dialog
          open={open}
          fullWidth
          maxWidth="md"
          BackdropProps={{
            style: { backgroundColor: 'rgba(0,0,30,0.4)' },
            invisible: true,
          }}
        >
          <DialogTitle fontWeight={500}>Cập nhật thông tin cơ bản</DialogTitle>
          <FormProvider methods={methods} onSubmit={handleSubmit(onSubmit)}>
            <Stack direction="row">
              {/* Left */}
              <Stack sx={{ px: 3, width: '60%', pr: 0 }}>
                {/* Basic infor */}
                <Stack justifyContent="space-between" direction="row" sx={{ pb: 2 }}>
                  <Select
                    isDisabled="true"
                    styles={customStyles}
                    name="prodId"
                    defaultValue={proId}
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
                    disabled={auction.winner_id > 0}
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
                        console.log(startDate.$d);
                      }}
                    />
                    <DateTimePicker
                      renderInput={(props) => <TextField {...props} sx={{ display: 'flex', flex: 1.5, ml: 1.5 }} />}
                      label="Ngày kết thúc"
                      value={endDate || dayjs()}
                      onChange={(newValue) => {
                        setEndDate(newValue);
                        console.log(endDate.$d);
                      }}
                    />
                  </LocalizationProvider>
                  {/* Active  */}
                  <FormControlLabel
                    sx={{ display: 'flex', flex: 1.2, fontSize: '0.2rem', ml: 1.5, mr: 0 }}
                    control={
                      <Switch
                        color="success"
                        disabled={auction.winner_id > 0}
                        checked={isActive}
                        onChange={(e) => setIsActive(e.target.checked)}
                      />
                    }
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
                disabled={auction.winner_id > 0 && auction.is_active === true}
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
        {/* Delete Product */}
        {!auction.is_default && auction.winner_id > 0 ? (
          <Tooltip enterDelay={1000} title="Phiên đấu giá đã có người tham gia">
            <span>
              <Button
                disabled={auction.winner_id > 0}
                key={auction.ID}
                onClick={handleClickOpenDelForm}
                sx={{
                  ml: 1,
                  borderRadius: 0,
                  p: 0,
                  bgcolor: 'transparent',
                  border: 'none',
                  fontSize: '0.9rem',
                  fontWeight: 400,
                  '&:hover': {
                    bgcolor: 'transparent',
                    textDecoration: 'underline',
                  },
                }}
              >
                Xoá
              </Button>
            </span>
          </Tooltip>
        ) : (
          <Button
            disabled={auction.winner_id > 0}
            key={auction.ID}
            onClick={handleClickOpenDelForm}
            sx={{
              ml: 1,
              borderRadius: 0,
              p: 0,
              bgcolor: 'transparent',
              border: 'none',
              fontSize: '0.9rem',
              fontWeight: 400,
              '&:hover': {
                bgcolor: 'transparent',
                textDecoration: 'underline',
              },
            }}
          >
            Xoá
          </Button>
        )}
        {/* Dialog Delete */}
        <Dialog
          sx={{ margin: 'auto', minWidth: '480px' }}
          BackdropProps={{
            style: { backgroundColor: 'rgba(0,0,30,0.2)' },
            invisible: true,
          }}
          fullScreen={fullScreen}
          open={openDelForm}
        >
          <Stack sx={{ p: 3 }}>Bạn có chắc muốn xoá phiên đấu giá này?</Stack>

          <Stack sx={{ p: 2 }} justifyContent="flex-end" direction="row" alignItems="center">
            <Button
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
              onClick={handleCloseDelForm}
            >
              Trở lại
            </Button>
            <Button
              key={auction.ID}
              color="error"
              variant="contained"
              sx={{
                ml: 1,
                width: '62px',
                color: 'white',
                bgcolor: '#F62217',
              }}
              onClick={() => {
                handleDelete(auction);
                handleCloseDelForm();
              }}
              autoFocus
            >
              Xoá
            </Button>
          </Stack>
        </Dialog>
      </Stack>
    </>
  );
}
