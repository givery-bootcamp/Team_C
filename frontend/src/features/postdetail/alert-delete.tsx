import { useRef, useState } from 'react';

import {
  AlertDialog,
  AlertDialogBody,
  AlertDialogContent,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogOverlay,
  Button,
  HStack,
  Input,
  Radio,
  RadioGroup,
  Text,
  useToast,
  VStack,
} from '@chakra-ui/react';

import { useNavigate } from 'react-router-dom';

import { useAppDispatch } from 'shared/hooks';

import { APIService } from 'shared/services';

interface PlayfulDeleteProps {
  isOpen: boolean;

  onClose: () => void;

  postId: number;
}

const PlayfulDelete = ({ isOpen, onClose, postId }: PlayfulDeleteProps) => {
  const [stage, setStage] = useState(0);
  const [riddleAnswer, setRiddleAnswer] = useState('');
  const [riddleLevel, setRiddleLevel] = useState('easy');
  const cancelRef = useRef<HTMLButtonElement>(null);
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const toast = useToast();
  interface Riddle {
    question: string;
    answer: string;
  }

  const getCurrentRiddle = (): Riddle => {
    const levelRiddles = riddles[riddleLevel];
    return levelRiddles[Math.floor(Math.random() * levelRiddles.length)];
  };

  const alerts = [
    '削除していいの？本当に削除していいの？',
    'さっき「削除」を押しましたよね？本当に本当に削除しちゃっていいんですか？',
    '削除したら泣いちゃうよ？？？',
    '本当に削除しちゃうの。。。？',
    'まだ間に合うよ？本当に削除する？',
    'ガチのまじの最後のチャンスだけど本当に削除します？',
  ];

  const riddles: { [key: string]: Riddle[] } = {
    easy: [
      {
        question: '私は頭も尻尾もありませんが、体はあります。私は何でしょう？',
        answer: '川',
      },
      { question: '丸いのに角がある。何でしょう？', answer: 'お饅頭' },
    ],
    medium: [
      {
        question:
          '私は常に先頭を歩きますが、決して後ろを向きません。私は何でしょう？',
        answer: '鼻',
      },
      {
        question: '重ければ重いほど、軽くなります。何でしょう？',
        answer: '風船',
      },
    ],
    hard: [
      {
        question:
          '私は昼も夜も動き続けますが、決して疲れません。私は何でしょう？',
        answer: '時計',
      },
      {
        question: '食べれば食べるほど大きくなります。何でしょう？',
        answer: '穴',
      },
    ],
  };

  const handleNextStage = () => {
    if (stage < alerts.length - 1) {
      setStage(stage + 1);
    } else {
      setStage(alerts.length);
    }
  };

  const handleDelete = async () => {
    const currentRiddle = getCurrentRiddle();
    if (riddleAnswer.toLowerCase() === currentRiddle.answer) {
      try {
        await dispatch(APIService.deletePost({ id: postId }));

        toast({
          title: '投稿が削除されました',

          description: 'あなたは賢明な選択をした',

          status: 'success',

          duration: 3000,

          isClosable: true,
        });

        navigate('/posts');
      } catch (error) {
        toast({
          title: 'エラー',

          description: '運命はあなたの投稿を守ったようだね。',

          status: 'error',

          duration: 3000,

          isClosable: true,
        });
      }
    } else {
      toast({
        title: '不正解',

        description: 'なぞなぞに正解できませんでした。投稿は安全です',

        status: 'warning',

        duration: 3000,

        isClosable: true,
      });
    }

    onClose();
    setStage(0);
    setRiddleAnswer('');
  };

  return (
    <AlertDialog
      isOpen={isOpen}
      leastDestructiveRef={cancelRef}
      onClose={onClose}
    >
      <AlertDialogOverlay>
        <AlertDialogContent>
          <AlertDialogHeader fontSize="lg" fontWeight="bold">
            {stage < alerts.length ? '削除の確認' : '最終試練'}
          </AlertDialogHeader>

          <AlertDialogBody>
            {stage < alerts.length ? (
              alerts[stage]
            ) : (
              <VStack spacing={4}>
                <Text>なぞなぞの難易度を選んでください：</Text>
                <RadioGroup onChange={setRiddleLevel} value={riddleLevel}>
                  <HStack spacing={4}>
                    <Radio value="easy">簡単</Radio>
                    <Radio value="medium">普通</Radio>
                    <Radio value="hard">難しい</Radio>
                  </HStack>
                </RadioGroup>
                <Text>{getCurrentRiddle().question}</Text>
                <Input
                  placeholder="答えを入力してください"
                  value={riddleAnswer}
                  onChange={(e) => setRiddleAnswer(e.target.value)}
                />
              </VStack>
            )}
          </AlertDialogBody>

          <AlertDialogFooter>
            <Button ref={cancelRef} onClick={onClose}>
              キャンセル
            </Button>

            {stage < alerts.length ? (
              <Button colorScheme="red" onClick={handleNextStage} ml={3}>
                はい、削除します
              </Button>
            ) : (
              <Button colorScheme="red" onClick={handleDelete} ml={3}>
                なぞなぞに答えて削除
              </Button>
            )}
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialogOverlay>
    </AlertDialog>
  );
};

export default PlayfulDelete;
