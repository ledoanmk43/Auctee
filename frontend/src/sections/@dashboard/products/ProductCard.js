import PropTypes from 'prop-types';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
// material
import { Box, Card, Link, Typography, Stack } from '@mui/material';
import { styled } from '@mui/material/styles';
// utils
import { fCurrency } from '../../../utils/formatNumber';
// components
import Label from '../../../components/Label';
import { ColorPreview } from '../../../components/color-utils';
import CountDown from '../../../utils/countdown';
// ----------------------------------------------------------------------

const ProductImgStyle = styled('img')({
  top: 0,
  width: '100%',
  height: '100%',
  objectFit: 'cover',
  position: 'absolute',
});

// ----------------------------------------------------------------------

ShopProductCard.propTypes = {
  product: PropTypes.object,
};

export default function ShopProductCard({ product }) {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate(`/auctee/auction/detail/?id=${product.Id}&product=${product.product_id}`);
  };
  return (
    <Stack
      justifyContent="space-between"
      sx={{ borderRadius: 0.3, minWidth: '172px', minHeight: '270px', display: 'flex', boxShadow: 3 }}
      cursor="pointer"
      onClick={handleClick}
    >
      <Box sx={{ pt: '100%', position: 'relative' }}>
        <ProductImgStyle
          alt={product.name}
          src={`/static/mock-images/products/product_${product.Id > 24 ? product.Id - 2 : product.Id}.jpg`}
        />
      </Box>

      <Stack justifyContent="space-between" spacing={2} sx={{ px: 1.5, py: 1, display: 'flex', minHeight: '98px' }}>
        <Typography
          sx={{
            overflow: 'hidden',
            textOverflow: 'ellipsis',
            display: '-webkit-box',
            WebkitLineClamp: '2',
            WebkitBoxOrient: 'vertical',
          }}
          variant="caption"
        >
          {product.name}
        </Typography>

        <Stack sx={{ mt: '5px !important' }} direction="row" alignItems="center" justifyContent="space-between">
          {/* <ColorPreview colors={product.product_id} /> */}
          <Stack variant="body2">
            <CountDown time={product.end_time} />
            {product.current_bid.toLocaleString('tr-TR', {
              style: 'currency',
              currency: 'VND',
            })}
          </Stack>
        </Stack>
      </Stack>
    </Stack>
  );
}
