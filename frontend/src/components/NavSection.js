import { useState } from 'react';
import PropTypes from 'prop-types';
import { NavLink as RouterLink, matchPath, useLocation } from 'react-router-dom';
// material
import { alpha, useTheme, styled } from '@mui/material/styles';
import { Box, List, Collapse, ListItemText, ListItemIcon, ListItemButton } from '@mui/material';
//
import Iconify from './Iconify';

// ----------------------------------------------------------------------

const ListItemStyle = styled((props) => <ListItemButton disableGutters {...props} />)(({ theme }) => ({
  ...theme.typography.body2,
  height: 40,
  position: 'relative',
  textTransform: 'capitalize',
  color: 'inherit',
  borderRadius: theme.shape.borderRadius,
}));

const ListItemIconStyle = styled(ListItemIcon)({
  color: 'inherit',
  display: 'flex',
  justifyContent: 'flex-end',
});

// ----------------------------------------------------------------------

NavItem.propTypes = {
  item: PropTypes.object,
  active: PropTypes.func,
};

function NavItem({ item, active }) {
  const theme = useTheme();

  const isActiveRoot = active(item.path);

  const { title, path, icon, info, children } = item;

  const [open, setOpen] = useState(isActiveRoot || true);

  const handleOpen = () => {
    setOpen((prev) => !prev);
  };

  const activeRootStyle = {
    color: 'primary.main',
    fontWeight: 'fontWeightMedium',
    bgcolor: alpha(theme.palette.primary.main, theme.palette.action.selectedOpacity),
  };

  const activeSubStyle = {
    color: theme.palette.background.main,
    fontWeight: 'fontWeightMedium',
  };

  if (children) {
    return (
      <>
        <ListItemStyle
          sx={{
            ...(isActiveRoot && activeRootStyle),
            '&:hover': { bgcolor: 'transparent' },
          }}
        >
          <ListItemIconStyle>{icon && icon}</ListItemIconStyle>
          <ListItemText sx={{ fontWeight: 500, ml: 0.5 }} disableTypography primary={title} />
          {info && info}
        </ListItemStyle>

        <Collapse in={open} timeout="auto" unmountOnExit>
          <List sx={{ mt: -2 }} component="div" disablePadding>
            {children.map((item) => {
              const { title, path } = item;
              const isActiveSub = active(path);

              return (
                <ListItemStyle
                  key={title}
                  component={RouterLink}
                  to={path}
                  sx={{
                    ...(isActiveSub && activeSubStyle),
                    ml: 3,
                    my: 1.5,
                    py: 0,
                    height: '20px',
                    '&:hover': { bgcolor: 'transparent', color: theme.palette.background.main },
                  }}
                >
                  <ListItemIconStyle />
                  <ListItemText disableTypography primary={title} />
                </ListItemStyle>
              );
            })}
          </List>
        </Collapse>
      </>
    );
  }

  return (
    <ListItemStyle
      component={RouterLink}
      to={path}
      sx={{
        ...(isActiveRoot && activeRootStyle),
      }}
    >
      <ListItemIconStyle>{icon && icon}</ListItemIconStyle>
      <ListItemText disableTypography primary={title} />
      {info && info}
    </ListItemStyle>
  );
}

NavSection.propTypes = {
  navConfig: PropTypes.array,
};

export default function NavSection({ navConfig, ...other }) {
  const { pathname } = useLocation();

  const match = (path) => (path ? !!matchPath({ path, end: false }, pathname) : false);

  return (
    <Box {...other}>
      <List disablePadding sx={{ p: 1 }}>
        {navConfig.map((item) => (
          <NavItem key={item.title} item={item} active={match} />
        ))}
      </List>
    </Box>
  );
}
