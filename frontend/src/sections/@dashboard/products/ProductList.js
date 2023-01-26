import PropTypes from 'prop-types';
// material
import { Grid } from '@mui/material';
import ShopProductCard from './ProductCard';

// ----------------------------------------------------------------------

ProductList.propTypes = {
  auctions: PropTypes.array.isRequired,
};

export default function ProductList({ auctions, ...other }) {
  return (
    <Grid container spacing={3} {...other}>
      {auctions.map((auction, index) => (
        <Grid key={index} item xs={12} sm={6} md={2}>
          <ShopProductCard auction={auction} />
        </Grid>
      ))}
    </Grid>
  );
}
