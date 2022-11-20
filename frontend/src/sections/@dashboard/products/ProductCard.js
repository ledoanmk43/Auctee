import PropTypes from 'prop-types';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
// material
import { Box, Card, Link, Typography, Stack } from '@mui/material';
import { styled } from '@mui/material/styles';
import Iconify from '../../../components/Iconify';

import CountDown from '../../../utils/countdown';
// ----------------------------------------------------------------------

const ProductImgStyle = styled('img')({
  top: 0,
  width: '100%',
  height: '100%',
  objectFit: 'cover',
  position: 'absolute',
});
const Keyframes = styled('div')({
  '@keyframes pull': {
    '0%': {
      transform: 'rotate(-70deg)',
    },
    '25%': {
      transform: 'rotate(-40deg)',
    },
    '50%': {
      transform: 'rotate(0deg)',
    },
    '75%': {
      transform: 'rotate(40deg)',
    },
    '100%': {
      transform: 'rotate(-70deg)',
    },
  },
  animation: 'pull 1s infinite ease-in-out',
  position: 'absolute',
  right: "-8px",
  top: '18px',
  transform: 'rotate(45deg)',
});

// ----------------------------------------------------------------------

ShopProductCard.propTypes = {
  auction: PropTypes.object,
};

export default function ShopProductCard({ auction }) {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate(`/auctee/auction/detail/?id=${auction.Id}&product=${auction.product_id}`);
  };
  return (
    <Stack
      justifyContent="space-between"
      sx={{
        borderRadius: 0.3,
        minWidth: '173px',
        minHeight: '270px',
        maxHeight: '270px',
        display: 'flex',
        boxShadow: 3,
      }}
      cursor="pointer"
      onClick={handleClick}
    >
      <Box sx={{ pt: '100%', position: 'relative' }}>
        <ProductImgStyle alt={auction.name} src={auction.image_path} />
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
          {auction.name}
        </Typography>

        <Stack
          sx={{ mt: '5px !important', position: 'relative' }}
          direction="row"
          alignItems="center"
          justifyContent="space-between"
        >
          {/* <ColorPreview colors={product.product_id} /> */}
          <Stack variant="body2">
            <CountDown text="Còn lại:" time={auction.end_time} />
            {auction.current_bid.toLocaleString('tr-TR', {
              style: 'currency',
              currency: 'VND',
            })}
            <Stack color="brown" direction="row">
              <Keyframes>
                <Iconify sx={{ transform: 'rotate(30deg)' }} icon={'mingcute:auction-line'} width={25} height={25} />
              </Keyframes>
            </Stack>
          </Stack>
        </Stack>
      </Stack>
    </Stack>
  );
}
