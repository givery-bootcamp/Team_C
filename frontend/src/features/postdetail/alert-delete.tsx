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

import { motion } from 'framer-motion';
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
  const [openHints, setOpenHints] = useState([false, false, false]);
  const [deleteChance, setDeleteChance] = useState(100);

  interface Riddle {
    question: string;
    answer: string[];
    hints: string[];
  }

  const handlehintClick = (index: number) => {
    const newOpenHints = [...openHints];
    newOpenHints[index] = !newOpenHints[index];
    setOpenHints(newOpenHints);

    const openCount = newOpenHints.filter((hint) => hint).length;
    let newDeleteChance;
    switch (openCount) {
      case 1:
        newDeleteChance = 90;
        break;
      case 2:
        newDeleteChance = 70;
        break;
      case 3:
        newDeleteChance = 50;
        break;
      default:
        newDeleteChance = 50;
    }
    setDeleteChance(newDeleteChance);
  };

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
          hints: [
            'アリは協力して何かをしているようです',
            '10匹のアリが何かを言っているようです',
            '言葉の中に「ありがとう」が含まれているかもしれません',
          ],
        },
        {
          question: '3缶（かん）にのったくだものは何ですか？',
          answer: ['みかん'],
          hints: [
            'アリは協力して何かをしているようです',
            '10匹のアリが何かを言っているようです',
            '言葉の中に「ありがとう」が含まれているかもしれません',
          ],
        },
      ],
      medium: [
        {
          question: '唐辛子（とうがらし）が怒られているよなにをしたのかな？',
          answer: ['からかった', '辛かった'],
          hints: [
            'アリは協力して何かをしているようです',
            '10匹のアリが何かを言っているようです',
            '言葉の中に「ありがとう」が含まれているかもしれません',
          ],
        },
        {
          question: '地面にある男の穴ってなぁに？',
          answer: ['マンホール'],
          hints: [
            'アリは協力して何かをしているようです',
            '10匹のアリが何かを言っているようです',
            '言葉の中に「ありがとう」が含まれているかもしれません',
          ],
        },
      ],
      hard: [
        {
          question: 'テレビやラジオにとりついているゆうれいってな～んだ？',
          answer: ['音量', '怨霊', 'おんりょう'],
          hints: [
            'アリは協力して何かをしているようです',
            '10匹のアリが何かを言っているようです',
            '言葉の中に「ありがとう」が含まれているかもしれません',
          ],
        },
        {
          question: '誉められたのって何年生？',
          answer: ['小学３年生', '小３', '賞賛'],
          hints: [
            'アリは協力して何かをしているようです',
            '10匹のアリが何かを言っているようです',
            '言葉の中に「ありがとう」が含まれているかもしれません',
          ],
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
    setOpenHints([false, false, false]);
    setDeleteChance(100);
  };

  useEffect(() => {
    if (stage === alerts.length) {
      const levelRiddles = riddles[riddleLevel];
      const random =
        levelRiddles[Math.floor(Math.random() * levelRiddles.length)];
      setCurrentRiddle(random);
    }
  }, [stage, riddleLevel, riddles, alerts]);

  const deleteButtonRef = useRef<HTMLButtonElement>(null);
  const [buttonPosition, setButtonPosition] = useState({
    x: window.innerWidth / 2,
    y: window.innerHeight / 2,
  });
  const [buttonScale, setButtonScale] = useState(1);
  const [buttonRotation, setButtonRotation] = useState(0);
  const circleRadius = 150;

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
    const handleResize = () => {
      const rect = deleteButtonRef.current?.getBoundingClientRect();
      if (rect) {
        setButtonPosition({
          x: window.innerWidth / 2 - rect.width / 2,
          y: window.innerHeight / 2 - rect.height / 2,
        });
      }
    };

    window.addEventListener('resize', handleResize);
    handleResize();

    return () => {
      window.removeEventListener('resize', handleResize);
    };
  }, []);

  const handleButtonClick = () => {
    handleDelete();
    setButtonScale(1.5);
    setButtonRotation(360);
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
                    {currentRiddle.hints.map((hint, i) => (
                      <HStack key={i}>
                        <Button
                          variant={openHints[i] ? 'solid' : 'ghost'}
                          onClick={() => handlehintClick(i)}
                        >
                          ヒント {i + 1}
                        </Button>
                        {openHints[i] && <Text>{hint}</Text>}
                      </HStack>
                    ))}
                    <Text>削除確率: {deleteChance}%</Text>
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
              <motion.div
                animate={{
                  x:
                    buttonPosition.x +
                    Math.cos((Date.now() / 1000) * Math.PI * 2) * circleRadius,
                  y:
                    buttonPosition.y +
                    Math.sin((Date.now() / 1000) * Math.PI * 2) * circleRadius,
                  scale: buttonScale,
                  rotate: buttonRotation,
                }}
                transition={{
                  duration: 5,
                  repeat: Infinity,
                  ease: 'linear',
                }}
              >
                <Button
                  ref={deleteButtonRef}
                  colorScheme="red"
                  onClick={handleButtonClick}
                  fontSize="2xl"
                  fontWeight="bold"
                  padding="1rem 2rem"
                  boxShadow="0px 0px 20px 5px rgba(255,0,0,0.5)"
                >
                  なぞなぞに答えて削除
                </Button>
              </motion.div>
            )}
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialogOverlay>
    </AlertDialog>
  );
};

export default PlayfulDelete;
