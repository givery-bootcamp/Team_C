import { useEffect, useMemo, useRef, useState } from 'react';

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
  const [currentRiddle, setCurrentRiddle] = useState<Riddle | null>(null);
  interface Riddle {
    question: string;
    answer: string[];
  }

  const alerts = useMemo(
    () => [
      '削除していいの？本当に削除していいの？',
      'さっき「削除」を押しましたよね？本当に本当に削除しちゃっていいんですか？',
      '削除したら泣いちゃうよ？？？',
      '本当に削除しちゃうの。。。？',
      'まだ間に合うよ？本当に削除する？',
      'ガチのまじの最後のチャンスだけど本当に削除します？',
    ],
    [],
  );

  const riddles: { [key: string]: Riddle[] } = useMemo(
    () => ({
      easy: [
        {
          question: 'アリが10匹で何かをいっていますが、その言葉は何ですか？',
          answer: ['ありがとう'],
        },
        {
          question: '3缶（かん）にのったくだものは何ですか？',
          answer: ['みかん'],
        },
      ],
      medium: [
        {
          question: '唐辛子（とうがらし）が怒られているよなにをしたのかな？',
          answer: ['からかった', '辛かった'],
        },
        {
          question: '地面にある男の穴ってなぁに？',
          answer: ['マンホール'],
        },
      ],
      hard: [
        {
          question: 'テレビやラジオにとりついているゆうれいってな～んだ？',
          answer: ['音量', '怨霊', 'おんりょう'],
        },
        {
          question: '誉められたのって何年生？',
          answer: ['小学３年生', '小３', '賞賛'],
        },
      ],
    }),
    [],
  );

  const handleRiddleLevelChange = (newLevel: string) => {
    setRiddleLevel(newLevel);
    const newLevelRiddles = riddles[newLevel];
    const random =
      newLevelRiddles[Math.floor(Math.random() * newLevelRiddles.length)];
    setCurrentRiddle(random);
    setRiddleAnswer('');
  };
  const handleNextStage = () => {
    if (stage < alerts.length - 1) {
      setStage(stage + 1);
    } else {
      setStage(alerts.length);
    }
  };

  const resetState = () => {
    setStage(0);
    setRiddleAnswer('');
    setCurrentRiddle(null);
  };

  const handleDelete = async () => {
    if (currentRiddle?.answer.includes(riddleAnswer)) {
      try {
        await dispatch(APIService.deletePost(postId));

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

  useEffect(() => {
    if (stage === alerts.length) {
      const levelRiddles = riddles[riddleLevel];
      const random =
        levelRiddles[Math.floor(Math.random() * levelRiddles.length)];
      setCurrentRiddle(random);
    }
  }, [stage, riddleLevel, riddles, alerts]);

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
                <RadioGroup
                  onChange={handleRiddleLevelChange}
                  value={riddleLevel}
                >
                  <HStack spacing={4}>
                    <Radio value="easy">簡単</Radio>
                    <Radio value="medium">普通</Radio>
                    <Radio value="hard">難しい</Radio>
                  </HStack>
                </RadioGroup>
                {currentRiddle && (
                  <>
                    <Text>{currentRiddle.question}</Text>
                    <Input
                      placeholder="答えを入力してください"
                      value={riddleAnswer}
                      onChange={(e) => setRiddleAnswer(e.target.value)}
                    />
                  </>
                )}
              </VStack>
            )}
          </AlertDialogBody>

          <AlertDialogFooter>
            <Button
              ref={cancelRef}
              onClick={() => {
                onClose();
                resetState();
              }}
            >
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
