import { useState, useEffect, lazy, Suspense, useContext } from 'react';
import { Link as RouterLink, useNavigate, useLocation, useSearchParams } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import Select from 'react-select';

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
  Avatar,
  Box,
  alertTitleClasses,
  FormControlLabel,
  IconButton,
} from '@mui/material';
import AddPhotoAlternateOutlinedIcon from '@mui/icons-material/AddPhotoAlternateOutlined';
import HighlightOffRoundedIcon from '@mui/icons-material/HighlightOffRounded';
import { LoadingButton } from '@mui/lab';
import { styled, useTheme } from '@mui/material/styles';
import { Icon } from '@iconify/react';
import { FormProvider, RHFTextField } from '../../components/hook-form';
import useLocationForm from '../update-address/useLocationForm';
import { ReloadContext } from '../../utils/Context';

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
  const [proQuantity, setQuantity] = useState('');
  const [expectPrice, setExpectPrice] = useState('');
  const [imageList, setImageList] = useState([]);
  const [images, setImages] = useState([]);
  const [optionList, setOptionList] = useState([]);
  const [optionColor, setOptionColor] = useState();
  const [optionSize, setOptionSize] = useState();
  const [optionModel, setOptionModel] = useState();

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
    const payload = {
      id: proId,
      name: proName,
      min_price: parseFloat(minPrice),
      expect_price: parseFloat(expectPrice),
      quantity: parseInt(proQuantity, 10),
      description: proDescription,
      product_images: images,
      product_options: optionList,
    };

    await fetch('http://localhost:8080/auctee/user/product', {
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
        setErrorMessage(res.message);
        setOpen(true);
      }
    });
  };

  const handleClickOpen = () => {
    setOpen(true);
  };
  const handleClickOpenCreateOption = () => {
    setOpenOption(true);
  };
  const [currentOption, setCurrentOption] = useState();
  const handleClickOpenUpdateOption = (option, i) => {
    setCurrentOption(i);
    setOptionColor(option.color);
    setOptionModel(option.model);
    setOptionSize(option.size);
    setOpenUpdateOption(true);
  };

  const [isLoading, setIsLoading] = useState(true);
  const handleUpdateAvatar = (e) => {
    setIsLoading(true);
    e.preventDefault();
    // const file = e.target.files[0];
    const filesArr = e.target.files;
    /* eslint-disable-next-line no-plusplus */
    for (let i = 0; i < filesArr.length; i++) {
      if (filesArr[i].size > 2000000) {
        alert('file too large');
        return;
      }
    }
    /* eslint-disable-next-line no-plusplus */
    for (let i = 0; i < filesArr.length; i++) {
      const reader = new FileReader();
      reader.readAsDataURL(filesArr[i]);
      reader.onload = () => {
        if (imageList?.length < 6) {
          imageList.unshift(reader.result); // base64encoded string
          images.unshift({ path: reader.result });
          setIsLoading(false);
        }
      };
      reader.onerror = (error) => {
        console.log('Error: ', error);
      };
    }
  };
  const handleDeleteImage = (index) => {
    setIsLoading(true);
    setImages(
      images.filter((item, i) => {
        return i !== index;
      })
    );
    setImageList(
      imageList.filter((item, i) => {
        return i !== index;
      })
    );
    setIsLoading(false);
  };

  const handleCreateOption = () => {
    const data = {
      color: optionColor,
      size: optionSize,
      model: optionModel,
    };
    optionList.push(data);
    setOptionColor('');
    setOptionModel('');
    setOptionSize('');
    setOpenOption(false);
  };
  const handleUpdateOption = (currentOption) => {
    optionList.map((option, i) => {
      if (i === currentOption) {
        option.color = optionColor;
        option.size = optionSize;
        option.model = optionModel;
      }
      setOptionColor('');
      setOptionModel('');
      setOptionSize('');
      return option;
    });
    setOpenUpdateOption(false);
  };
  const handleDeleteOption = (index) => {
    setOptionList(
      optionList.filter((item, i) => {
        return i !== index;
      })
    );
  };
  useEffect(() => {
    setIsReloading(false);
  }, [isFetching, isReloading]);

  return (
    <>
      <Stack justifyContent="space-between" alignItems="center" direction="row" sx={{ maxHeight: '100%', pb: 2 }}>
        <Typography fontSize={'1.2rem'} variant="body2" sx={{ color: 'black' }}>
          Tất cả sản phẩm
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
          &nbsp; Thêm sản phẩm mới
        </Button>
      </Stack>

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
              {/* Images */}
              <Stack direction="row" sx={{ flex: 1.5 }}>
                <Stack sx={{ mb: 3 }} alignItems="center" direction="row">
                  {imageList.map((item, index) => (
                    <Stack key={index} sx={{ position: 'relative', mx: 1 }}>
                      <Avatar
                        variant="square"
                        alt={item.toString()}
                        src={item.toString()}
                        sx={{ width: 60, height: 60, borderRadius: 0.5 }}
                      />
                      <IconButton
                        onClick={() => handleDeleteImage(index)}
                        sx={{
                          maxHeight: '15px',
                          maxWidth: '15px',
                          position: 'absolute',
                          top: '-6px',
                          right: '-6px',
                          p: 0,
                          bgcolor: 'transparent',
                          color: 'red',
                          '&:hover': { color: 'white', bgcolor: 'red' },
                        }}
                      >
                        <HighlightOffRoundedIcon sx={{ fontSize: '1.2rem' }} />
                      </IconButton>
                    </Stack>
                  ))}
                  {/* button */}
                  <Button
                    disabled={imageList?.length > 5}
                    sx={{
                      textTransform: 'none',
                      borderRadius: 0.5,
                      border: '1px dashed #bdbdbd',
                      my: 1,
                      mx: 1,
                      px: '20px !important',
                      width: 60,
                      height: 60,
                      color: theme.palette.background.main,
                      bgcolor: 'white',
                      display: 'flex',
                      flexDirection: 'column',
                      justifyContent: 'space-around',
                      '&:hover': {
                        opacity: 0.75,
                      },
                    }}
                    variant="square"
                    color="error"
                    component="label"
                  >
                    <AddPhotoAlternateOutlinedIcon />
                    <Typography variant="caption" sx={{ width: '50px', fontSize: '0.6rem', textAlign: 'center' }}>
                      {`Thêm hình ảnh (${imageList.length}/6)`}
                    </Typography>
                    <input onChange={(e) => handleUpdateAvatar(e)} hidden accept="image/*" multiple type="file" />
                  </Button>
                  {imageList?.length < 5 && (
                    <Stack sx={{ mx: 1 }}>
                      <Typography variant="caption" sx={{ color: 'black' }}>
                        Dụng lượng file tối đa 1 MB,
                      </Typography>
                      <Typography variant="caption" sx={{ color: 'black' }}>
                        Định dạng:.JPEG, .PNG
                      </Typography>
                    </Stack>
                  )}
                </Stack>
              </Stack>
              {/* Basic infor */}
              <Stack justifyContent="space-between" direction="row" sx={{ pb: 2 }}>
                <RHFTextField
                  color="error"
                  required
                  label="Mã sp"
                  name="productid"
                  type="text"
                  value={proId || ''}
                  onChange={(e) => setProId(e.target.value)}
                  size="small"
                  variant="outlined"
                  sx={{ width: '22%' }}
                  InputLabelProps={{ fontSize: '0.5rem' }}
                />
                <RHFTextField
                  color="error"
                  required
                  label="Tên sản phẩm"
                  name="name"
                  type="text"
                  value={proName || ''}
                  onChange={(e) => setProName(e.target.value)}
                  size="small"
                  variant="outlined"
                  sx={{ width: '73%' }}
                />
              </Stack>
              <Stack justifyContent="space-between" alignItems="center" direction="row" sx={{ pb: 2 }}>
                <RHFTextField
                  min="1"
                  color="error"
                  required
                  label="Số lượng"
                  name="quantity"
                  type="number"
                  value={proQuantity || ''}
                  onChange={(e) => setQuantity(e.target.value)}
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
                  onChange={(e) => setMinPrice(e.target.value)}
                  size="small"
                  variant="outlined"
                  sx={{ width: '27%' }}
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
                  onChange={(e) => setExpectPrice(e.target.value)}
                  size="small"
                  variant="outlined"
                  sx={{ width: '27%' }}
                />
                <Typography variant="body1">VNĐ</Typography>
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
                />
              </Stack>
              <Stack alignItems="center" direction="row" sx={{ height: 3 }}>
                {error && (
                  <Typography variant="body2" color="error">
                    Mã hàng đã tồn tại
                  </Typography>
                )}
              </Stack>
            </Stack>
            {/* Right */}
            <Stack direction="row" sx={{ px: 3, width: '37.2%' }}>
              <Divider orientation="vertical" />
              <Stack alignItems="stretch" direction="column" sx={{ ml: 3, pr: 0, minWidth: '100%', mr: 3 }}>
                {/* Button add option */}
                <Stack direction="row" alignItems="center" justifyContent="space-between" sx={{ mb: 2 }}>
                  <Typography variant="body1" sx={{ color: 'black' }}>
                    Phân loại
                  </Typography>
                  <Button
                    disableRipple
                    onClick={handleClickOpenCreateOption}
                    sx={{
                      px: '15px !important',
                      color: 'white',
                      bgcolor: theme.palette.background.main,
                      fontWeight: 500,
                      textTransform: 'none',
                    }}
                    variant="contained"
                    color="error"
                    component="label"
                  >
                    <Icon icon="akar-icons:plus" />
                    &nbsp; Thêm tuỳ chọn
                  </Button>
                </Stack>
                {/* Option list */}
                <Stack
                  sx={{
                    overflowX: 'hidden',
                    overflowY: 'auto',
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
                    width: '100%',
                    maxHeight: '290px',
                    bgcolor: 'white',
                  }}
                >
                  {optionList.map((option, index) => (
                    <Stack key={index}>
                      {/* options */}
                      <Stack
                        alignItems="center"
                        direction="row"
                        sx={{
                          overflow: 'hidden',
                          mr: 1,
                          mb: 1,
                          width: '100%',
                          minHeight: '10%',
                        }}
                      >
                        <Typography
                          key={index}
                          variant="subtitle2"
                          onClick={() => handleClickOpenUpdateOption(option, index)}
                          sx={{
                            // bgcolor: `${openUpdateOption && index ? theme.palette.background.main : 'white'}`,
                            px: 1,
                            border: `1px solid ${theme.palette.background.main}`,
                            borderRadius: 0.4,
                            '&:hover': {
                              bgcolor: theme.palette.background.main,
                              boxShadow: 2,
                              color: 'white',
                            },
                            overflow: 'hidden',
                            textOverflow: 'ellipsis',
                            display: '-webkit-box',
                            WebkitLineClamp: '1',
                            WebkitBoxOrient: 'vertical',
                          }}
                        >
                          {option.model} - {option.color} - size: &nbsp;
                          {option.size}
                        </Typography>
                        <IconButton
                          onClick={() => handleDeleteOption(index)}
                          sx={{
                            ml: 1,
                            maxHeight: '20px',
                            maxWidth: '20px',
                            p: 0,
                            bgcolor: 'transparent',
                            color: 'red',
                            '&:hover': { color: 'white', bgcolor: 'red' },
                          }}
                        >
                          <HighlightOffRoundedIcon sx={{ fontSize: '1.6rem', fontWeight: 400 }} />
                        </IconButton>
                      </Stack>
                      {/* Form update option */}
                      {openUpdateOption && (
                        <Dialog
                          open={openUpdateOption}
                          maxWidth="400px"
                          BackdropProps={{
                            style: { backgroundColor: 'rgba(0,0,30,0.1)' },
                            invisible: true,
                          }}
                        >
                          <DialogTitle fontWeight={500}>Chỉnh sửa tuỳ chọn</DialogTitle>
                          <Stack justifyContent="space-between" sx={{ pb: 2, px: 3, minHeight: '170px' }}>
                            <RHFTextField
                              color="error"
                              required
                              label="Màu sắc"
                              name="color"
                              type="text"
                              value={optionColor || ''}
                              onChange={(e) => setOptionColor(e.target.value)}
                              size="small"
                              variant="outlined"
                              InputLabelProps={{ fontSize: '0.5rem' }}
                            />
                            <RHFTextField
                              color="error"
                              required
                              label="Size"
                              name="size"
                              type="text"
                              value={optionSize || ''}
                              onChange={(e) => setOptionSize(e.target.value)}
                              size="small"
                              variant="outlined"
                              InputLabelProps={{ fontSize: '0.5rem' }}
                            />
                            <RHFTextField
                              color="error"
                              required
                              label="Brand"
                              name="model"
                              type="text"
                              value={optionModel || ''}
                              onChange={(e) => setOptionModel(e.target.value)}
                              size="small"
                              variant="outlined"
                              InputLabelProps={{ fontSize: '0.5rem' }}
                            />
                          </Stack>
                          <Stack
                            justifyContent="center"
                            alignItems="center"
                            direction="row"
                            sx={{ mt: 2, pb: 4, position: 'relative' }}
                          >
                            <Button
                              size="medium"
                              variant="outlined"
                              onClick={() => {
                                setOptionColor('');
                                setOptionModel('');
                                setOptionSize('');
                                setOpenUpdateOption(false);
                              }}
                              sx={{
                                px: 1.6,
                                position: 'absolute',
                                right: 114,
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
                            <Button
                              disableRipple
                              color="error"
                              sx={{ px: 3, position: 'absolute', right: 24 }}
                              size="medium"
                              variant="contained"
                              onClick={() => handleUpdateOption(currentOption)}
                            >
                              Lưu
                            </Button>
                          </Stack>
                        </Dialog>
                      )}
                    </Stack>
                  ))}
                </Stack>
              </Stack>
              {/* Form create option */}
              {openCreateOption && (
                <Dialog
                  open={openCreateOption}
                  maxWidth="400px"
                  BackdropProps={{
                    style: { backgroundColor: 'rgba(0,0,30,0.4)' },
                    invisible: true,
                  }}
                >
                  <DialogTitle fontWeight={500}>Thêm tuỳ chọn</DialogTitle>
                  <Stack justifyContent="space-between" sx={{ pb: 2, px: 3, minHeight: '170px' }}>
                    <RHFTextField
                      color="error"
                      required
                      label="Màu sắc"
                      name="color"
                      type="text"
                      value={optionColor || ''}
                      onChange={(e) => setOptionColor(e.target.value)}
                      size="small"
                      variant="outlined"
                      InputLabelProps={{ fontSize: '0.5rem' }}
                    />
                    <RHFTextField
                      color="error"
                      required
                      label="Size"
                      name="size"
                      type="text"
                      value={optionSize || ''}
                      onChange={(e) => setOptionSize(e.target.value)}
                      size="small"
                      variant="outlined"
                      InputLabelProps={{ fontSize: '0.5rem' }}
                    />
                    <RHFTextField
                      color="error"
                      required
                      label="Brand"
                      name="model"
                      type="text"
                      value={optionModel || ''}
                      onChange={(e) => setOptionModel(e.target.value)}
                      size="small"
                      variant="outlined"
                      InputLabelProps={{ fontSize: '0.5rem' }}
                    />
                  </Stack>
                  <Stack
                    justifyContent="center"
                    alignItems="center"
                    direction="row"
                    sx={{ mt: 2, pb: 4, position: 'relative' }}
                  >
                    <Button
                      size="medium"
                      variant="outlined"
                      onClick={() => setOpenOption(false)}
                      sx={{
                        px: 1.6,
                        position: 'absolute',
                        right: 114,
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
                    <Button
                      disableRipple
                      color="error"
                      sx={{ px: 2, position: 'absolute', right: 24 }}
                      size="medium"
                      variant="contained"
                      onClick={() => handleCreateOption()}
                    >
                      Thêm
                    </Button>
                  </Stack>
                </Dialog>
              )}
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
              sx={{ px: 3, position: 'absolute', right: 24 }}
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
