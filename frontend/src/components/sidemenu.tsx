import { Box, VStack } from "@chakra-ui/react";
import { NavLink } from "react-router-dom";

const SideMenu = () => {
    return (
        <VStack p={3}>
            <Box>
                <NavLink to="/posts" className="app-menu__nav-item">
                    トップ
                </NavLink>
            </Box>
        </VStack>
    );
}

export default SideMenu;