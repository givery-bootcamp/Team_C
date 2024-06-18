import { Flex, VStack, Link } from "@chakra-ui/react";

const SideMenu = () => {
    return (
        <VStack w={'100px'} paddingTop={3}>
            <Flex w={'100%'} >
                <Link href="/posts" color={'blue.500'}>
                    トップへ
                </Link>
            </Flex>
        </VStack>
    );
}
export default SideMenu;